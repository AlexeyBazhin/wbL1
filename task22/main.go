package main

import (
	"fmt"
	"strconv"
	"strings"
)

type BigInt []int

func main() {
	str1 := "111"
	str2 := "999999"
	//идея - хранить большие числа поразрядно в массиве, так для 1.000.000 len(bigInt) = 7 = {0,0,0,0,0,0,1}
	big1 := NewBigInt(str1)
	big2 := NewBigInt(str2)
	//fmt.Println(big2)
	//fmt.Println(big1.IsGreater(big2))
	fmt.Println(big1.Add(big2))
	fmt.Println(big2.Sub(big1))
	fmt.Println(big2.Multiple(big1))

}

// конструктор BigInt - bigInt[0] - младший разряд - последний символ строки
func NewBigInt(bigIntStr string) BigInt {
	bigInt := make([]int, len(bigIntStr))
	for i, j := len(bigInt)-1, 0; i >= 0; i-- {
		bigInt[i], _ = strconv.Atoi(bigIntStr[j : j+1])
		j++
	}
	return bigInt
}

// превращаем из массива в строку {0,0,0,0,0,0,1} -> 1000000
func (bigInt BigInt) String() string {
	builder := &strings.Builder{}
	for i := len(bigInt) - 1; i >= 0; i-- {
		builder.WriteString(strconv.Itoa(bigInt[i]))
	}
	return builder.String()
}

// суммирование столбиком
func (big1 BigInt) Add(big2 BigInt) BigInt {
	//сверху всегда большее число,
	if !big1.IsGreater(big2) {
		big1, big2 = big2, big1
	}
	result := make([]int, len(big1))
	buff := 0                        //суммирование происходит поразрядно, хранит переполнение разряда
	for i := 0; i < len(big1); i++ { //итерируемся по разрядам длинного числа
		var sum int //сумма соответствующих разрядов
		if i < len(big2) {
			sum = big1[i] + big2[i] + buff //если у короткого числа есть соответствующий разряд
		} else {
			sum = big1[i] + buff //если нет
		}
		result[i] = sum % 10 //пример 9 + 9 = 18, записываем 8
		buff = sum / 10      //1 отдаем буферу, который на следующей итерации прибавит к сумме переполнение этой итерации
	}

	if buff != 0 { // 99 + 1: result = {0,0} buf = 1
		result = append(result, buff) //result = {0,0,1}
	}

	return result
}

func (big1 BigInt) Sub(big2 BigInt) BigInt {
	//операции над отрицательными числами не рализованы, поэтому всегда вчитаем меньшее из большего
	if !big1.IsGreater(big2) {
		big1, big2 = big2, big1
		fmt.Println("Второе число больше первого! Вычитание первого из второго: ")
	}

	var result BigInt
	var borrow int
	for i := 0; i < len(big1); i++ {
		var diff int

		if i < len(big2) {
			diff = big1[i] - big2[i] - borrow //не забываем вычитать разряд, из которого заняли
		} else {
			diff = big1[i] - borrow
		}

		//13 - 5: 3 - 5 = - 8 нужно прибавить 10 и запомнить, что заняли из старшего разряда
		if diff < 0 {
			diff += 10
			borrow = 1
		} else {
			borrow = 0 //обнуляем заем, если разница разрядов с учетом заема >0
		}

		result = append(result, diff)
	}

	return result
}

func (big1 BigInt) Multiple(big2 BigInt) BigInt {

	if !big1.IsGreater(big2) {
		big1, big2 = big2, big1
	}
	//если меньшее число = 0
	if big2.String() == "0" {
		return big2
	}

	//умножение столбиком предполагает суммирование результатов поразрядного умножения
	//123 * 34 = 492 + 3690
	var sums []BigInt
	for i := 0; i < len(big2); i++ {

		var digitsMuls []string //123 * 4 = 12 + 80 + 400
		for j := 0; j < len(big1); j++ {
			//умножаем разряды, например 2 * 4 = 8
			digitsMulStr := strconv.Itoa(big1[j] * big2[i])
			//переводим к строке, чтобы приписать нужное количество 0: 8 -> 80
			for k := 0; k < j; k++ {
				digitsMulStr = strings.Join([]string{digitsMulStr, "0"}, "")
			}
			digitsMuls = append(digitsMuls, digitsMulStr)
		}

		sum := BigInt{0}
		//у нас хранятся строки 12 80 400. Суммируем их
		for j := 0; j < len(digitsMuls); j++ {
			sum = sum.Add(NewBigInt(digitsMuls[j]))
		}
		//дальше приписываем нужное количество 0
		//для начального примера 123 * 34 = 123 * 4 + 123 * 30, но мы делаем 123 * 3, поэтому взависимости от i (итератор по 34) мы приписываем 0
		//
		sumStr := sum.String()
		for j := 0; j < i; j++ {
			sumStr = strings.Join([]string{sumStr, "0"}, "")
		}
		sums = append(sums, NewBigInt(sumStr))
	}

	//скалдываем все результаты умножений верхнего числа на каждый разряд
	result := BigInt{0}
	for _, sum := range sums {
		result = result.Add(sum)
	}

	return result
}

func (firstBig BigInt) IsGreater(secondBig BigInt) bool {
	if firstLen, secondLen := len(firstBig), len(secondBig); firstLen == secondLen {
		for i := firstLen - 1; i >= 0; i-- {
			if firstBig[i] > secondBig[i] {
				return true
			} else if firstBig[i] < secondBig[i] {
				return false
			}
		}
		return false
	}
	if len(firstBig) < len(secondBig) {
		return false
	}
	return true
}
