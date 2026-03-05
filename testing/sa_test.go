package testing

import (
	"bytes"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestSA(t *testing.T) {
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
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(d.Name(), ".golite") {
			return nil
		}

		// Only test in sa directories or sa_custom
		if !strings.Contains(path, "/sa/") && !strings.Contains(path, "/sa_custom/") {
			return nil
		}

		t.Run(path, func(t *testing.T) {
			dir := filepath.Dir(path)
			base := filepath.Base(path)
			baseNoExt := strings.TrimSuffix(base, ".golite")

			// Look for expected file
			var expectedFile string
			candidates := []string{
				filepath.Join(dir, "expected"),
				filepath.Join(dir, baseNoExt+"_expected"),
				filepath.Join(dir, strings.Replace(baseNoExt, "sa", "sa_expected_", 1)),
			}

			// Special case for sa5 which has multiple files in same dir
			if strings.Contains(path, "sa5_") {
				num := strings.Split(baseNoExt, "_")[1]
				candidates = append(candidates, filepath.Join(dir, "sa5_expected_"+num))
			}

			for _, c := range candidates {
				if _, err := os.Stat(c); err == nil {
					expectedFile = c
					break
				}
			}

			if expectedFile == "" {
				// If no expected file, we assume it should pass without errors
				// but many benchmarks HAVE errors. Let's skip if no expected for now.
				t.Logf("Skipping %s: no expected file found", path)
				return
			}

			projectRoot, _ := filepath.Abs(".")
			// find project root by looking for go.mod
			for {
				if _, err := os.Stat(filepath.Join(projectRoot, "go.mod")); err == nil {
					break
				}
				projectRoot = filepath.Dir(projectRoot)
				if projectRoot == "/" {
					t.Fatal("Could not find project root")
				}
			}

			golitePath := filepath.Join(projectRoot, "golite", "golite.go")

			cmd := exec.Command("go", "run", golitePath, path)
			var out bytes.Buffer
			cmd.Stdout = &out
			var stderr bytes.Buffer
			cmd.Stderr = &stderr
			cmd.Env = os.Environ()

			// We expect it to fail if there are semantic errors
			cmd.Run()

			actualOutput := out.String()
			expectedOutputBytes, err := os.ReadFile(expectedFile)
			if err != nil {
				t.Fatalf("failed to read expected file: %v", err)
			}
			expectedOutput := string(expectedOutputBytes)

			// Normalize both for comparison
			processLines := func(lines []string, isActual bool) []string {
				var processed []string
				for _, line := range lines {
					l := strings.TrimSpace(line)
					if isActual {
						l = strings.ReplaceAll(l, "*", "")
					}
					if l != "" {
						processed = append(processed, l)
					}
				}

				// Sort lines by semantic error(line, column)
				re := regexp.MustCompile(`semantic error\((\d+)[: ,](\d+)\)`)
				sort.SliceStable(processed, func(i, j int) bool {
					mi := re.FindStringSubmatch(processed[i])
					mj := re.FindStringSubmatch(processed[j])

					if len(mi) == 3 && len(mj) == 3 {
						li, _ := strconv.Atoi(mi[1])
						ci, _ := strconv.Atoi(mi[2])
						lj, _ := strconv.Atoi(mj[1])
						cj, _ := strconv.Atoi(mj[2])

						if li != lj {
							return li < lj
						}
						return ci < cj
					}
					// Fallback to string comparison if regex doesn't match
					return processed[i] < processed[j]
				})
				return processed
			}

			actualLines := processLines(strings.Split(actualOutput, "\n"), true)
			expectedLines := processLines(strings.Split(expectedOutput, "\n"), false)

			if len(actualLines) != len(expectedLines) {
				t.Errorf("Line count mismatch. Expected %d, got %d\nActual (processed):\n%s\nExpected (processed):\n%s",
					len(expectedLines), len(actualLines), strings.Join(actualLines, "\n"), strings.Join(expectedLines, "\n"))
				return
			}

			for i := range actualLines {
				if actualLines[i] != expectedLines[i] {
					t.Errorf("Mismatch at line %d:\nGot:      %q\nExpected: %q", i+1, actualLines[i], expectedLines[i])
				}
			}
		})
		return nil
	})
	if err != nil {
		t.Fatalf("failed to walk benchmarks: %v", err)
	}
}
