package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}

	//первое решение
	sequentialSum(getSquareChan(arr, wg))
	//второе решение
	concurrentSum(getSquareChan(arr, wg))
	//третье решение
	concurrentSumAtomic(getSquareChan(arr, wg))
}

// лучший вариант получения квадратов из предыдущего задания
func getSquareChan(arr []int, wg *sync.WaitGroup) chan int {
	squareChan := make(chan int)
	for _, elem := range arr {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			squareChan <- number * number
		}(elem)
	}

	//Главная фишка - ожидание WG, после чего закрытие канала в отдельной горутине.
	go func() {
		wg.Wait()
		close(squareChan) //Закрываем канал, чтобы можно было использовать range и не зависить по размеру от других программных сущностей.
	}()
	return squareChan
}
func sequentialSum(squareChan chan int) {
	var sum int
	for square := range squareChan {
		sum += square //пока в канал поступают значения, прибавляем их последовательно к сумме
	}
	fmt.Println(sum)
}
func concurrentSum(squareChan chan int) {
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	var sum int
	for square := range squareChan {
		wg.Add(1)
		//конкурентно прибавляем к sum
		go func(mu *sync.Mutex, square int) {
			defer wg.Done()
			mu.Lock() //прибавление к сумме происходит в отдельных горутинах, sum - общий разделяемый ресурс.
			//defer mu.Unlock() //отложенные функции хранятся с использованием стека. Сначала выполнится данная строчка, потом wg.Done().
			sum += square // благодаря мьютексу можем избавиться от race condition
			mu.Unlock()
		}(mu, square)
	}
	wg.Wait()
	mu.Lock()
	fmt.Println(sum)
	mu.Unlock()
}
func concurrentSumAtomic(squareChan chan int) {
	wg := &sync.WaitGroup{}
	var sum int64 //atomic.AddInt64 требует int64
	for square := range squareChan {
		wg.Add(1)
		//конкурентно прибавляем к sum
		go func(square int) {
			defer wg.Done()
			atomic.AddInt64(&sum, int64(square)) //атомарно инкрементируем сумму
		}(square)
	}
	wg.Wait()
	fmt.Println(sum)
}
