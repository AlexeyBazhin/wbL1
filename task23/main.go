package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(removeIndex(arr, 0))
	fmt.Printf("Оригинальный: %v\n",arr)
	fmt.Println(removeLoop(arr, 0))
	fmt.Printf("Оригинальный: %v\n",arr)	
	fmt.Println(removeCopy(arr, 0))
	fmt.Printf("Оригинальный: %v\n",arr)
}

func removeIndex(arr []int, index int) ([]int, error) {
	if index < 0 || index > len(arr) {
		return nil, fmt.Errorf("index %v out of range %v", index, len(arr))
	}
	var newArr []int
	newArr = append(newArr, arr[:index]...)
	newArr = append(newArr, arr[index+1:]...)	
	return newArr, nil
	//на первый взгляд хочется сделать так:
	//arr = append(arr[:index], arr[index+1:]...)
	//но даже при том, что сам arr теперь будет ссылаться на другой массив,
	//внутри append первый параметр, arr[:index], поэтому изменится оригинальный массив
}

// помещаем удаляемый элемент в конец массива, передвигая идущие после него
// после чего возвращаем массив без последнего элемента
func removeLoop(arr []int, index int) ([]int, error) {
	if index < 0 || index > len(arr) {
		return nil, fmt.Errorf("index %v out of range %v", index, len(arr))
	}
	//данный способ будет изменять оригинальный массив
	// for i := index; i < len(arr)-1; i++ {
	// 	arr[i] = arr[i+1]
	// }
	// return arr[:len(arr)-1], nil

	//но можно сделать так:
	// var newArr []int
	// newArr = append(newArr, arr...)
	// for i := index; i < len(newArr)-1; i++ {
	// 	newArr[i] = newArr[i+1]
	// }
	// return newArr[:len(arr)-1], nil

	//либо проще
	newArr := make([]int, len(arr)-1)
	for i := index; i < len(newArr); i++ {
		newArr[i] = arr[i+1]
	}
	return newArr, nil
}
func removeCopy(arr []int, index int) ([]int, error) {
	if index < 0 || index > len(arr) {
		return nil, fmt.Errorf("index %v out of range %v", index, len(arr))
	}
	//данный способ тоже изменит оригинальный массив
	//copy(arr[index:], arr[index+1:])
	//return arr[:len(arr)-1], nil

	//делаем так:
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	copy(newArr[index:], newArr[index+1:])
	return newArr[:len(arr)-1], nil
}
