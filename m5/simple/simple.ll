source_filename = "m5/simple/simple"
target triple = "arm64-apple-macosx14.0.0"


declare ptr @malloc(i32)
declare i32 @scanf(ptr, ...)
declare i32 @printf(ptr, ...)
declare void @free(ptr)
%struct.Point2D = type {i64, i64}

@globalInit = global i64 0
@.fmt0 = private unnamed_addr constant [4 x i8] c"%ld\00"
@.fmt1 = private unnamed_addr constant [29 x i8] c"offset=%d\0Apt2.x=%d\0Apt2.y=%d\0A\00"

define ptr @Init(i64 %initVal) {
entry:
  %t0 = alloca i64
  store i64 %initVal, ptr %t0
  %t1 = alloca ptr
  store ptr null, ptr %t1
  %t2 = load i64, ptr %t0
  %t3 = icmp sgt i64 %t2, 0
  br i1 %t3, label %L0, label %L1
L0:
  %t4 = call ptr @malloc(i32 16)
  %t5 = bitcast ptr %t4 to ptr
  store ptr %t5, ptr %t1
  %t6 = load i64, ptr %t0
  %t7 = load ptr, ptr %t1
  %t8 = getelementptr %struct.Point2D, ptr %t7, i32 0, i32 0
  store i64 %t6, ptr %t8
  %t9 = load i64, ptr %t0
  %t10 = load ptr, ptr %t1
  %t11 = getelementptr %struct.Point2D, ptr %t10, i32 0, i32 1
  store i64 %t9, ptr %t11
  %t12 = load ptr, ptr %t1
  ret ptr %t12
L1:
  br label %L2
L2:
  %t13 = load ptr, ptr %t1
  ret ptr %t13
}

define void @main() {
entry:
  %t0 = alloca i64
  %t1 = alloca i64
  %t2 = alloca ptr
  %t3 = alloca ptr
  store i64 5, ptr %t0
  %t4 = load i64, ptr %t0
  %t5 = add i64 %t4, 7
  %t6 = mul i64 %t5, 3
  store i64 %t6, ptr %t1
  %t7 = call ptr @malloc(i32 16)
  %t8 = bitcast ptr %t7 to ptr
  store ptr %t8, ptr %t2
  %t9 = load i64, ptr %t0
  %t10 = load ptr, ptr %t2
  %t11 = getelementptr %struct.Point2D, ptr %t10, i32 0, i32 0
  store i64 %t9, ptr %t11
  %t12 = load i64, ptr %t1
  %t13 = load ptr, ptr %t2
  %t14 = getelementptr %struct.Point2D, ptr %t13, i32 0, i32 1
  store i64 %t12, ptr %t14
  %t15 = getelementptr [4 x i8], ptr @.fmt0, i32 0, i32 0
  call i32 (i8*, ...) @scanf(ptr %t15, ptr @globalInit)
  %t16 = load i64, ptr @globalInit
  %t17 = call ptr @Init(i64 %t16)
  store ptr %t17, ptr %t3
  %t18 = getelementptr [29 x i8], ptr @.fmt1, i32 0, i32 0
  %t19 = load i64, ptr @globalInit
  %t20 = load ptr, ptr %t3
  %t21 = getelementptr %struct.Point2D, ptr %t20, i32 0, i32 0
  %t22 = load i64, ptr %t21
  %t23 = load ptr, ptr %t3
  %t24 = getelementptr %struct.Point2D, ptr %t23, i32 0, i32 1
  %t25 = load i64, ptr %t24
  call i32 (i8*, ...) @printf(ptr %t18, i64 %t19, i64 %t22, i64 %t25)
  %t26 = load ptr, ptr %t2
  %t27 = bitcast ptr %t26 to ptr
  call void @free(ptr %t27)
  %t28 = load ptr, ptr %t3
  %t29 = bitcast ptr %t28 to ptr
  call void @free(ptr %t29)
  ret void
}
