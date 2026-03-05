package testing

import (
	"bytes"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func normalize(s string) []string {
	var preliminaryLines []string
	for _, line := range strings.Split(s, "\n") {
		// Strip comments
		if idx := strings.Index(line, "//"); idx != -1 {
			line = line[:idx]
		}

		// Inject spaces around punctuation (including parentheses) to treat them as individual tokens
		punctuation := []string{"(", ")", "{", "}", "[", "]", ",", ";", "*", "+", "-", "/", "=", "<", ">", "!", " ", "\n"}
		for _, p := range punctuation {
			// line = strings.ReplaceAll(line, p, " "+p+" ")
			line = strings.ReplaceAll(line, p, "")
		}

		// Collapse whitespace
		fields := strings.Fields(line)
		if len(fields) > 0 {
			preliminaryLines = append(preliminaryLines, strings.Join(fields, " "))
		}
	}

	// Pull up braces '{' or keywords 'else' on newlines to the previous line
	var lines []string
	for i := 0; i < len(preliminaryLines); i++ {
		current := preliminaryLines[i]
		// Pull up '{' or 'else ...' if preceded by a line (usually ending in '}')
		if len(lines) > 0 && (current == "{" || strings.HasPrefix(current, "else")) {
			lines[len(lines)-1] = lines[len(lines)-1] + " " + current
		} else {
			lines = append(lines, current)
		}
	}

	return lines
}

func compare(t *testing.T, path string, actual, expected []string) {
	minLen := len(actual)
	if len(expected) < minLen {
		minLen = len(expected)
	}

	for i := 0; i < minLen; i++ {
		if actual[i] != expected[i] {
			t.Errorf("Diff at line %d:\nExpected: %q\nGot:      %q", i+1, expected[i], actual[i])
		}
	}

	if len(actual) != len(expected) {
		t.Errorf("AST output mismatch for %s: line count mismatch (got %d, expected %d)", path, len(actual), len(expected))
	}
}

func TestAST(t *testing.T) {
	// Find the project root (where benchmarks is)
	root := "../benchmarks"
	if _, err := os.Stat(root); os.IsNotExist(err) {
		root = "benchmarks"
		if _, err := os.Stat(root); os.IsNotExist(err) {
			t.Fatal("Could not find benchmarks directory")
		}
	}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".golite") {
			t.Run(path, func(t *testing.T) {
				projectRoot, _ := filepath.Abs(filepath.Dir(root))
				if filepath.Base(root) == "benchmarks" {
					projectRoot = filepath.Dir(root)
				}
				golitePath := filepath.Join(projectRoot, "golite", "golite.go")

				cmd := exec.Command("go", "run", golitePath, "-ast", path)
				var out bytes.Buffer
				cmd.Stdout = &out
				var stderr bytes.Buffer
				cmd.Stderr = &stderr
				cmd.Env = os.Environ()

				if err := cmd.Run(); err != nil {
					t.Fatalf("compiler failed: %v\nStderr: %s", err, stderr.String())
				}

				original, err := os.ReadFile(path)
				if err != nil {
					t.Fatalf("failed to read file: %v", err)
				}

				compare(t, path, normalize(out.String()), normalize(string(original)))
			})
		}
		return nil
	})
	if err != nil {
		t.Fatalf("failed to walk benchmarks: %v", err)
	}
}
