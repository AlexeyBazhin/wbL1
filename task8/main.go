package main

import "fmt"

func main() {
	number := 25
	fmt.Printf("bin: %b\n", number)
	number = setOne(number, 2)
	fmt.Printf("bin: %b; dec: %d\n", number, number)

	number = setZero(number, 2)
	fmt.Printf("bin: %b; dec: %d\n", number, number)

	number = invertBit(number, 2)
	fmt.Printf("bin: %b; dec: %d\n", number, number)

	number = invertBit(number, 2)
	fmt.Printf("bin: %b; dec: %d\n", number, number)
}

// нужно выполнить дизъюнкцию с числом, в котором только i-й бит установлен в 1
func setOne(number int, index int) int {
	mask := (1 << index)
	fmt.Printf("mask: bin: %b; dec: %d\n", mask, mask)
	return mask | number //можно было сразу return number | (1<< index)
}

// нужно выполнить конъюнкцию с числом, в котором только i-й бит отличается от изменяемого числа
func setZero(number int, index int) int {
	//для этого выполняем xor между оригинальным числом и числом, в котором только i-й бит установлен в 1
	//в других битах единицы останутся единицами (1 xor 0 = 1), а нули нулями (0 xor 0 = 0)
	//изменяемый бит, если был равен 1, то в маске станет равным 0
	mask := number ^ (1 << index)
	fmt.Printf("mask: bin: %b; dec: %d\n", mask, mask)
	//но, если он был равен нулю, то в маске он станет равным 1
	//пример: 1101 xor 10 = 1111
	//поэтому нужно производить конъюнкцию, чтоб, в случае изменения 0 бита на 0, число оставалось тем же
	return mask & number
	//либо же делать конъюнкцию только в случае,когда маска стала больше оригинального числа
	//if mask > number {
	//return mask & number
	//}
	//return mask
}

// в случае, когда нужно именно инвертировать определенный бит, достаточно xor
func invertBit(number int, index int) int {
	return number ^ (1 << index)
}
