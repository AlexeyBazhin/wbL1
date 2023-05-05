package main

import "fmt"

func main() {

	a, b := 10, 21
	b, a = a, b //встроенная возможность go
	fmt.Println(a, b)

	// 2 способ - арифметические операции
	b = a - b
	a = a - b // a - (a-b) = b
	b = a + b // b + (a-b) = a
	fmt.Println(a, b)

	// 2.5 способ
	b = b + a
	a = b - a // (b + a) - a = b
	b = b - a // (b + a) - b = a
	fmt.Println(a, b)
	// 2,75 способ)))
	a = a * b
	b = a / b // (a * b) / b = a
	a = a / b // (a * b) / a = b
	fmt.Println(a, b)

	// 3 способ - операция xor
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)

}
