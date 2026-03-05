# proj
The starter code for the compiler project

# Build Instructions

```bash
$ source build.sh
```

# Usage Instructions

```bash
# printing all the tokens parsed
$ ./glc -l filename.golite
# printing the ast
$ ./glc -ast filename.golite
# output stack-based LLVM IR instead of register-based SSA
$ ./glc -llvm-stack filename.golite
# compile LLVM to ARM64 assembly
$ ./glc -target="aarch64-linux-gnu" -S filename.golite
# specifying the target llvm string
$ ./glc -target=STRING filename.golite
```

# Testing Instructions

```bash
$ go test -v ./testing/... 2>&1 | grep -E "AST output mismatch|Diff at line|Expected:|Got:"
$ go test -v ./testing/ast_test.go
$ go test -v ./testing/sa_test.go
```

# Cluster Build & Execution Instructions

1. **Build Compiler and Generate Assembly (e.g., on a `jammy` node):**
   Use the `cluster_build.sh` script to build the compiler and generate the ARM64 assembly files (`.s`).
   ```bash
   $ source cluster_build.sh
   
   # Or manually generate assembly for a specific file:
   $ ./glc -target="aarch64-linux-gnu" -S benchmarks/arm/arm.golite
   ```

2. **Compile and Execute Assembly (on the `wing` server):**
   Once the `.s` file is generated, compile it with `gcc` and run the resulting binary.
   ```bash
   $ gcc benchmarks/arm/arm.s -o benchmarks/arm/arm_bin && ./benchmarks/arm/arm_bin
   ```
