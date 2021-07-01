// func VecSum(a []int64) int64
TEXT ·VecSum(SB),$64-32 //загружаем 3 переменные длинной по 64 бита
  MOVQ a+0(FP), SI              // address of a
  MOVQ len+8(FP), CX    // array length


  VPXOR Y4,Y4, Y4 // обнуляем векторный регистр
  MOVQ $0, DX     
  CMPQ CX,$0
  JE end

start:
  VMOVDQU (SI), Y1    //заносим значения массива в векторный регистр
  VMOVDQU (SI), Y2
  VPMULDQ Y1, Y2, Y3  //перемноножем векторные регистры 
  VPADDQ Y3,Y4,Y4 //складываем в счетчик
  ADDQ $32, SI
  SUBQ $4, CX
  JNLE start

  VMOVDQU Y4, d0-32(SP)//заносим занчение векторного регистра суммы в стек
  LEAQ d0-32(SP), BX   // заносим адрес стека в BX
  MOVQ $4, AX          // array length

sum:  //{sum1(int64),sum2(int64),sum3(int64)...} переводим в sum(int64)
  ADDQ (BX), DX
  ADDQ $8,  BX
  DECQ AX
  JNZ  sum
  
end:
  MOVL DX, ret+24(FP)   // final result
  RET
