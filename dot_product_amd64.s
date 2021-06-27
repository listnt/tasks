// definition of func VDotProdAVX2(a[] int32, b[] int32) int32
TEXT ·VDotProdAVX2(SB), $64-56
  MOVQ a+0(FP), AX    // address of a
  MOVQ b+24(FP), BX   // address of b
  MOVQ len+8(FP), CX  // array length
  VPXORD Y4, Y4, Y4
start:
  VMOVDQU32 (AX), Y1
  VMOVDQU32 (BX), Y2
  VPMULLD Y1, Y2, Y3
  VPADDD Y3, Y4, Y4
  ADDQ $32, AX
  ADDQ $32, BX
  SUBQ $8, CX
  JNZ start
  VMOVDQU32 Y4, d0-32(SP)

  LEAQ d0-32(SP), BX
  MOVQ $8, AX //array length
  MOVQ $0, CX //sum of bx vec
sum:
  ADDL (BX), CX
  ADDQ $4,  BX
  DECQ AX
  JNZ  sum
  MOVL CX, AX
  MOVL AX, ret+48(FP)
  RET
