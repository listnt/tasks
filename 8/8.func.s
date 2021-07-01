// func BitMask(a int64,b int64,c bool) int64
TEXT ·BitMask(SB),$0-32 //загружаем 3 переменные длинной по 64 бита
	MOVQ a+0(FP), AX    //Заносим переменную в регистр AX
	MOVQ b+8(FP), BX    //Заносим переменную в регистр BX
    MOVQ c+16(FP),DX    //Заносим переменную в регистр DX, 
    //Из-за выравнивания переменная занимает 8 байт x86-x64 
    CMPQ DX, $0         //Проверяем на равность нулю
    JE  bitr            //Если равно переходим к метке bitr(bit reser)
    //Если не равно переходим к метке bits(bit set)
bits:                   
    BTSQ BX,AX          //Устанавливаем в регистре AX бит BX = 1
    JMP end             //Прыгаем к метке конец
bitr:                   //Метка BITR
    BTRQ BX,AX          //Сбрасываем бит BX в регистре AX в 0
end:                    //Метка end
	MOVQ AX, ret+24(FP) //Заносим результат в возвращаемое значение
    RET
    