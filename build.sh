echo "generating the antlr lexer and parser"
cd grammars
source generate.sh
cd ..
echo "finished generating the antlr lexer and parser"

echo "building the golite binary"
go build -o glc golite/golite.go
echo "successfully built the golite binary"

# ./glc benchmarks/bad.golite
# ./glc -l -ast benchmarks/sa/sa1/sa1.golite
# ./glc benchmarks/sa/sa1/sa1.golite
# ./glc benchmarks/sa/sa2/sa2.golite
# ./glc benchmarks/sa/sa3/sa3.golite
# ./glc benchmarks/sa/sa4/sa4.golite
# ./glc benchmarks/sa/sa5/sa5_1.golite
# ./glc benchmarks/sa/sa5/sa5_2.golite
# ./glc benchmarks/sa/sa5/sa5_3.golite
# ./glc benchmarks/sa/sa5/sa5_4.golite
# ./glc benchmarks/sa/sa6/sa6.golite
# ./glc benchmarks/sa/sa7/sa7.golite
# ./glc benchmarks/sa/sa8/sa8.golite
# ./glc benchmarks/sa/sa9/sa9.golite
# ./glc benchmarks/sa/sa10/sa10.golite
# ./glc benchmarks/sa/sa11/sa11.golite
# ./glc benchmarks/sa/sa13/sa13.golite
# ./glc benchmarks/sa/sa14/sa14.golite
# ./glc benchmarks/sa/sa15/sa15.golite

echo "Testing the llvm stack benchmarks"
diff -b benchmarks/binarytree/output <(./glc -target="arm64-apple-macosx14.0.0" -llvm-stack benchmarks/binarytree/binarytree.golite && lli benchmarks/binarytree/binarytree.ll < benchmarks/binarytree/input)
diff -b benchmarks/hard/output <(./glc -target="arm64-apple-macosx14.0.0" -llvm-stack benchmarks/hard/hard.golite && lli benchmarks/hard/hard.ll < benchmarks/hard/input)
diff -b benchmarks/linkedlist/output <(./glc -target="arm64-apple-macosx14.0.0" -llvm-stack benchmarks/linkedlist/linkedlist.golite && lli benchmarks/linkedlist/linkedlist.ll < benchmarks/linkedlist/input)
diff -b benchmarks/mixed/output <(./glc -target="arm64-apple-macosx14.0.0" -llvm-stack benchmarks/mixed/mixed.golite && lli benchmarks/mixed/mixed.ll < benchmarks/mixed/input)
diff -b benchmarks/powmod/output <(./glc -target="arm64-apple-macosx14.0.0" -llvm-stack benchmarks/powmod/powmod.golite && lli benchmarks/powmod/powmod.ll < benchmarks/powmod/input)
diff -b benchmarks/primes/output <(./glc -target="arm64-apple-macosx14.0.0" -llvm-stack benchmarks/primes/primes.golite && lli benchmarks/primes/primes.ll < benchmarks/primes/input)
diff -b benchmarks/primes2/output <(./glc -target="arm64-apple-macosx14.0.0" -llvm-stack benchmarks/primes2/primes2.golite && lli benchmarks/primes2/primes2.ll < benchmarks/primes2/input)
diff -b benchmarks/thermopylae/output <(./glc -target="arm64-apple-macosx14.0.0" -llvm-stack benchmarks/thermopylae/thermopylae.golite && lli benchmarks/thermopylae/thermopylae.ll < benchmarks/thermopylae/input)
diff -b benchmarks/Twiddleedee/output <(./glc -target="arm64-apple-macosx14.0.0" -llvm-stack benchmarks/Twiddleedee/Twiddleedee.golite && lli benchmarks/Twiddleedee/Twiddleedee.ll < benchmarks/Twiddleedee/input)
diff -b benchmarks/arm/output <(./glc -target="arm64-apple-macosx14.0.0" -llvm-stack benchmarks/arm/arm.golite && lli benchmarks/arm/arm.ll < benchmarks/arm/input)
echo "Finished testing the llvm stack benchmarks"

echo "Testing the ssa llvm benchmarks"
diff -b benchmarks/arm/output <(./glc -target="arm64-apple-macosx14.0.0" -S benchmarks/arm/arm.golite && gcc-15 benchmarks/arm/arm.s -o benchmarks/arm/arm_bin && ./benchmarks/arm/arm_bin < benchmarks/arm/input)
# diff -b benchmarks/binarytree/output <(./glc -target="arm64-apple-macosx14.0.0" -S benchmarks/binarytree/binarytree.golite && gcc-15 benchmarks/binarytree/binarytree.s -o benchmarks/binarytree/binarytree_bin && ./benchmarks/binarytree/binarytree_bin < benchmarks/binarytree/input)
# diff -b benchmarks/hard/output <(./glc -target="arm64-apple-macosx14.0.0" -S benchmarks/hard/hard.golite && gcc-15 benchmarks/hard/hard.s -o benchmarks/hard/hard_bin && ./benchmarks/hard/hard_bin < benchmarks/hard/input)
# diff -b benchmarks/linkedlist/output <(./glc -target="arm64-apple-macosx14.0.0" -S benchmarks/linkedlist/linkedlist.golite && gcc-15 benchmarks/linkedlist/linkedlist.s -o benchmarks/linkedlist/linkedlist_bin && ./benchmarks/linkedlist/linkedlist_bin < benchmarks/linkedlist/input)
# diff -b benchmarks/mixed/output <(./glc -target="arm64-apple-macosx14.0.0" -S benchmarks/mixed/mixed.golite && gcc-15 benchmarks/mixed/mixed.s -o benchmarks/mixed/mixed_bin && ./benchmarks/mixed/mixed_bin < benchmarks/mixed/input)
# diff -b benchmarks/powmod/output <(./glc -target="arm64-apple-macosx14.0.0" -S benchmarks/powmod/powmod.golite && gcc-15 benchmarks/powmod/powmod.s -o benchmarks/powmod/powmod_bin && ./benchmarks/powmod/powmod_bin < benchmarks/powmod/input)
# diff -b benchmarks/primes/output <(./glc -target="arm64-apple-macosx14.0.0" -S benchmarks/primes/primes.golite && gcc-15 benchmarks/primes/primes.s -o benchmarks/primes/primes_bin && ./benchmarks/primes/primes_bin < benchmarks/primes/input)
# diff -b benchmarks/primes2/output <(./glc -target="arm64-apple-macosx14.0.0" -S benchmarks/primes2/primes2.golite && gcc-15 benchmarks/primes2/primes2.s -o benchmarks/primes2/primes2_bin && ./benchmarks/primes2/primes2_bin < benchmarks/primes2/input)
# diff -b benchmarks/thermopylae/output <(./glc -target="arm64-apple-macosx14.0.0" -S benchmarks/thermopylae/thermopylae.golite && gcc-15 benchmarks/thermopylae/thermopylae.s -o benchmarks/thermopylae/thermopylae_bin && ./benchmarks/thermopylae/thermopylae_bin < benchmarks/thermopylae/input)
# diff -b benchmarks/Twiddleedee/output <(./glc -target="arm64-apple-macosx14.0.0" -S benchmarks/Twiddleedee/Twiddleedee.golite && gcc-15 benchmarks/Twiddleedee/Twiddleedee.s -o benchmarks/Twiddleedee/Twiddleedee_bin && ./benchmarks/Twiddleedee/Twiddleedee_bin < benchmarks/Twiddleedee/input)
echo "Finished testing the ssa llvm benchmarks"

# go run golite/golite.go -l benchmarks/sa/sa1/sa1.golite 

# rm glc
