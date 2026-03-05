package testing

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestLLVMSSA(t *testing.T) {
	root := "../benchmarks"
	if _, err := os.Stat(root); os.IsNotExist(err) {
		root = "benchmarks"
	}

	projectRoot, _ := filepath.Abs(filepath.Dir(root))
	if filepath.Base(root) == "benchmarks" {
		projectRoot = filepath.Dir(root)
	}
	golitePath := filepath.Join(projectRoot, "golite", "golite.go")

	entries, err := os.ReadDir(root)
	if err != nil {
		t.Fatalf("failed to read benchmarks dir: %v", err)
	}

	for _, d := range entries {
		if !d.IsDir() {
			continue
		}

		benchName := d.Name()
		benchDir := filepath.Join(root, benchName)
		goliteFile := filepath.Join(benchDir, benchName+".golite")
		inputFile := filepath.Join(benchDir, "input")
		outputFile := filepath.Join(benchDir, "output")

		if _, err := os.Stat(goliteFile); os.IsNotExist(err) {
			continue
		}

		t.Run(benchName, func(t *testing.T) {
			tmpGoliteFile := filepath.Join("/tmp", benchName+".golite")
			tmpLLFile := filepath.Join("/tmp", benchName+".ll")

			inputData, err := os.ReadFile(goliteFile)
			if err != nil {
				t.Fatalf("failed to read source file: %v", err)
			}
			if err := os.WriteFile(tmpGoliteFile, inputData, 0644); err != nil {
				t.Fatalf("failed to write tmp file: %v", err)
			}

			cmd := exec.Command("go", "run", golitePath, "-target=arm64-apple-macosx14.0.0", tmpGoliteFile)
			var stderr bytes.Buffer
			cmd.Stderr = &stderr
			if err := cmd.Run(); err != nil {
				t.Fatalf("compiler failed: %v\nStderr: %s", err, stderr.String())
			}

			defer os.Remove(tmpGoliteFile)
			defer os.Remove(tmpLLFile)

			expectedOutput := ""
			hasExpected := false
			if outData, err := os.ReadFile(outputFile); err == nil {
				expectedOutput = string(outData)
				hasExpected = true
			}

			lliCmd := exec.Command("lli", tmpLLFile)
			if _, err := os.Stat(inputFile); err == nil {
				inData, _ := os.ReadFile(inputFile)
				lliCmd.Stdin = bytes.NewReader(inData)
			}

			var lliOut bytes.Buffer
			var lliErr bytes.Buffer
			lliCmd.Stdout = &lliOut
			lliCmd.Stderr = &lliErr

			// Ignore the run error because lli exits with the return value of main(),
			// which is often non-zero for these benchmarks!
			_ = lliCmd.Run()

			if !hasExpected {
				return
			}

			actual := normalizeOutput(lliOut.String())
			expected := normalizeOutput(expectedOutput)

			if actual != expected {
				t.Errorf("Output mismatch for %s\nExpected:\n%s\n\nGot:\n%s\nStderr: %s", benchName, expected, actual, lliErr.String())
			}
		})
	}
}
