// func VecSum(a []int64) int64
TEXT ·VecSum(SB),$64-32 //загружаем 3 переменные длинной по 64 бита
  MOVQ a+0(FP), SI              // address of a
  MOVQ len+8(FP), CX    // array length


  VPXOR Y4,Y4, Y4 // обнуляем векторный регистр
  MOVQ $0, R11
  CMPQ CX,$4
  JL vector_loop_end
vector_loop_begin:
  VMOVDQU (SI),Y1
  VMOVDQU (SI),Y2
  VPMULDQ Y1,Y2,Y3
  VPADDQ Y3,Y4,Y4
  ADDQ $32,SI
  SUBQ $4,CX
  CMPQ CX, $4
  JGE vector_loop_begin

  VMOVDQU Y4, d0-32(SP)//заносим занчение векторного регистра суммы в стек
  LEAQ d0-32(SP), BX   // заносим адрес стека в BX

  ADDQ (BX),R11
  ADDQ $8,  BX

  ADDQ (BX),R11
  ADDQ $8,  BX
  
  ADDQ (BX),R11
  ADDQ $8,  BX
  
  ADDQ (BX),R11
  ADDQ $8,  BX

vector_loop_end:
  CMPQ CX,$0
  JLE scalar_loop_end
scalar_loop_begin:
  MOVQ (SI),AX
  MULQ AX
  ADDQ AX, R11
  ADDQ $8,SI
  DECQ CX
  CMPQ CX,$0
  JG scalar_loop_begin
scalar_loop_end:
end: 
  MOVQ R11, ret+24(FP)   // final result
  RET
