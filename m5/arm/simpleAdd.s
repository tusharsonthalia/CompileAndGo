        .arch armv8-a 
        .text 

        .type main, %function 
        .global main 
        .p2align 2 

main: 
    //main code
    sub sp, sp, #32
    //// Prologue
    sub sp, sp, 16
    stp x29, x30, [sp]
    mov x29, sp //assume a = [sp + 32] b = [sp + 24] c = [sp + 16]
    ///////////
    mov x0, #129
    str x0, [sp, #32]
    adrp x0, .READ
    add    x0, x0, :lo12:.READ
    add x1, sp, #24
    bl scanf
    ldr x2, [sp, #32]
    ldr x3, [sp, #24]
    add x4, x2, x3
    str x4, [sp, #16]
    adrp x0, .PRINT_LN
    add     x0, x0, :lo12:.PRINT_LN
    ldr x1, [sp, #32]
    bl printf
    adrp x0, .PRINT_LN
    add     x0, x0, :lo12:.PRINT_LN
    ldr x1, [sp, #24]
    bl printf
    adrp x0, .PRINT
    add     x0, x0, :lo12:.PRINT
    ldr x1, [sp, #16]
    bl printf
    /// Epilogue
    ldp x29, x30, [sp], #16
    add sp, sp, #32
    ret
    /////////
    .size main, (. - main)

.READ:
	.asciz	"%ld"
	.size	.READ, 4

.PRINT:
	.asciz	"%ld"
	.size	.PRINT, 4

.PRINT_LN:
	.asciz	"%ld\n"
	.size	.PRINT_LN, 5
