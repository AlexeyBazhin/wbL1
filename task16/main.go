package main

import "fmt"

func main() {
	arr := []int{10,12,36,15,4,12,5,18,20}
	quickSort(arr, 0, len(arr)-1)
	fmt.Printf("Отсортированный массив: %v\n", arr)
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(arr, left, right) //получаем индекс точки опоры - все элементы слева от нее (включая) меньше всех элементов справа
	fmt.Printf("Массив:%v, left:%v, right:%v, pivot:%v\n", arr, left, right, pivot)
	quickSort(arr, left, pivot) //сортируем левую часть
	quickSort(arr, pivot+1, right) //сортируем правую часть
}
func partition(arr []int, i, j int) int {
	pivot := arr[(i+j)/2] //берем точкой опоры элемент по середине
	
	for i < j { //пока указатели  не сойдутся
		for arr[i] < pivot { //ищем элементы, которые больше опоры и находятся слева
			i++
		}
		for arr[j] > pivot { //ищем элементы, которые меньше опоры и находятся справа
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i] //меняем местами элемент больший опоры и меньший
		i++
		j--
	}
	return j
}
