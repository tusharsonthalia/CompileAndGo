source_filename = "simple"
target triple = "arm64-apple-macosx14.0.0"
%struct.Point2D = type {i64, i64}
@globalInit = common global i64 0

define %struct.Point2D* @Init(i64 %initVal)
{
L0:
	%_ret_val = alloca %struct.Point2D*
	%_P_initVal = alloca i64
	%newPt = alloca %struct.Point2D*
	store i64 %initVal, i64* %_P_initVal
	store ptr null, %struct.Point2D** %newPt
	%r0 = load i64, i64* %_P_initVal
	%r2 = icmp sgt i64 %r0, 0
	br i1 %r2, label %L2, label %L3
L2:
	%r3 = call i8* @malloc(i32 16)
	%r4 = bitcast i8* %r3 to %struct.Point2D*
	store %struct.Point2D* %r4, %struct.Point2D** %newPt
	%r5 = load i64, i64* %_P_initVal
	%r6 = load %struct.Point2D*, %struct.Point2D** %newPt
	%r7 = getelementptr %struct.Point2D, %struct.Point2D* %r6, i32 0, i32 0
	store i64 %r5, i64* %r7
	%r8 = load i64, i64* %_P_initVal
	%r9 = load %struct.Point2D*, %struct.Point2D** %newPt
	%r10 = getelementptr %struct.Point2D, %struct.Point2D* %r9, i32 0, i32 1
	store i64 %r8, i64* %r10
	%r11 = load %struct.Point2D*, %struct.Point2D** %newPt
	store %struct.Point2D* %r11, %struct.Point2D** %_ret_val
	br label %L4
L3:
	br label %L4
L4:
	%r12 = load %struct.Point2D*, %struct.Point2D** %newPt
	store %struct.Point2D* %r12, %struct.Point2D** %_ret_val
	br label %L1
L1:
	%r13 = load %struct.Point2D*, %struct.Point2D** %_ret_val
	ret %struct.Point2D* %r13

}
define i64 @main()
{
L5:
	%_ret_val = alloca i64
	%a = alloca i64
	%b = alloca i64
	%pt1 = alloca %struct.Point2D*
	%pt2 = alloca %struct.Point2D*
	store i64 5, i64* %a
	%r14 = load i64, i64* %a
	%r15 = add i64 %r14, 7
	%r16 = mul i64 %r15, 3
	store i64 %r16, i64* %b
	%r17 = call i8* @malloc(i32 16)
	%r18 = bitcast i8* %r17 to %struct.Point2D*
	store %struct.Point2D* %r18, %struct.Point2D** %pt1
	%r19 = load i64, i64* %a
	%r20 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r21 = getelementptr %struct.Point2D, %struct.Point2D* %r20, i32 0, i32 0
	store i64 %r19, i64* %r21
	%r22 = load i64, i64* %b
	%r23 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r24 = getelementptr %struct.Point2D, %struct.Point2D* %r23, i32 0, i32 1
	store i64 %r22, i64* %r24
	call i8 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0),i64* @globalInit)
	%r25 = load i64, i64* @globalInit
	%r26 = call i64  @Init(i64 %r25)
	store i64 %r26, %struct.Point2D** %pt2
	%r27 = load i64, i64* @globalInit
	%r28 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r29 = getelementptr %struct.Point2D, %struct.Point2D* %r28, i32 0, i32 0
	%r30 = load i64, i64* %r29
	%r31 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r32 = getelementptr %struct.Point2D, %struct.Point2D* %r31, i32 0, i32 1
	%r33 = load i64, i64* %r32
	call i8 (i8*, ...) @printf(i8* getelementptr inbounds ([32 x i8], [32 x i8]* @.fstr0, i32 0, i32 0),i64 %r27,i64 %r30,i64 %r33)
	%r34 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r35 = bitcast %struct.Point2D* %r34 to i8*
	call void @free(i8* %r35)
	%r36 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r37 = bitcast %struct.Point2D* %r36 to i8*
	call void @free(i8* %r37)
	store i64 0, i64* %_ret_val
	br label %L6
L6:
	%r38 = load i64, i64* %_ret_val
	ret i64 %r38

}

declare i8* @malloc(i32)
declare i32 @scanf(i8*, ...)
declare i32 @printf(i8*, ...)
declare void @free(i8*)
@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr0 = private unnamed_addr constant [32 x i8] c"offset=%ld\0Apt2.x=%ld\0Apt2.y=%ld\0A\00", align 1