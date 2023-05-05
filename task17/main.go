package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []int{1, 2, 2, 2, 2, 2, 3, 3, 5}
	//arr := []int{2}
	fmt.Printf("Первый найденый индекс с рекурсией: %v: %v\n", arr, firstFoundBinarySearchRecur(arr, 2))
	fmt.Printf("Первый найденый индекс с циклом: %v: %v\n", arr, firstFoundBinarySearchLoop(arr, 2))
	fmt.Printf("Левый индекс с циклом: %v: %v\n", arr, leftBinarySearch(arr, 2))
	fmt.Printf("Правый индекс с циклом: %v: %v\n", arr, rightBinarySearch(arr, 2))
}

// Находит первый совпадающий элемент, используя рекурсивный подход
// найденный индекс не обязательно будет левым, правым или средним вхождением искомого числа, если в массиве есть несколько вхождений этого числа
func firstFoundBinarySearchRecur(arr []int, number int) int {
	left, right := 0, len(arr)-1
	mid := (left + right) / 2
	if arr[mid] == number { //совпадает - возвращаем
		return mid
	}
	if left == right { //значит пришел массив длины 1, его элемент уже проверился и не совпадает, значит искомого элемента нет
		return -1
	}
	if arr[mid] > number {
		return firstFoundBinarySearchRecur(arr[:mid], number) //значит искомое число содержится в левой половине массива
	} else {
		return firstFoundBinarySearchRecur(arr[mid+1:], number) //значит искомое число содержится в правой половине массива
	}
}

// Находит первый совпадающий элемент, используя рекурсивный подход
// использование цикла будет предпочтительнее рекурсии, потому что зачем нагружать стек вызовов и выделять лишнюю память, когда можно этого не делать
func firstFoundBinarySearchLoop(arr []int, number int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == number {
			return mid
		}
		if arr[mid] > number {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// поиск левого индекса вхождения элемента
func leftBinarySearch(arr []int, number int) int {
	left, right := 0, len(arr)-1
	for left != right { //пока не останется 1 элемент
		mid := (left + right) / 2
		if arr[mid] < number {
			left = mid + 1 //елси проверяемое число меньше - проверяем правую половину
		} else {
			right = mid //т.к. мы не знаем данный элемент больше или равен, то проверяем левую половину, включая этот элемент
		}
	}
	if arr[left] == number { //без разницы, left = right
		return left
	} else {
		return -1
	}
}

func rightBinarySearch(arr []int, number int) int {
	left, right := 0, len(arr)-1
	for left != right { //пока не останется 1 элемент
		mid := int(math.Ceil(float64(left+right) / 2)) //здесь важно проводить округление вверх, т.к. в случае [2,2] из-за неопределенности <= зациклимся
		if arr[mid] > number {
			right = mid - 1 //елси проверяемое число больше - проверяем левуб половину
		} else {
			left = mid //т.к. мы не знаем данный элемент меньше или равен, то проверяем правую половину, включая этот элемент
		}
	}
	if arr[right] == number { //без разницы, left = right
		return right
	} else {
		return -1
	}
}
