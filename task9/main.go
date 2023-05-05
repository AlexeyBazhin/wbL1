package main

import (
	"fmt"
	"time"
)

func main() {
	generated()
	//manual()
}

func generated(){
	xChan := make(chan int)
	x2Chan := make(chan int)

	go func() {
		counter := 0
		for {			
			select {
			case <-time.After(1 * time.Microsecond): 
				close(xChan) //закрываем канал, чтоб горутина х2 знала, что конвейр остановился
				return
			default: //пока не сработает таймер записываем данные
				xChan <- counter
				counter++
			}
		}
	}()

	go func() {
		//закрываем канал, после того как считаем все данные из xChan и отправим их в x2Chan
		defer close(x2Chan)
		for x := range xChan {
			x2Chan <- x * 2
		}
	}()

	//считываем все данные x * 2
	for x2 := range x2Chan {
		fmt.Println(x2)
	}
}
func manual() {
	input := []int{1, 2, 3, 4, 5}

	xChan := make(chan int)
	x2Chan := make(chan int)

	go func() {
		defer close(xChan)
		for _, x := range input {
			xChan <- x
		}
	}()

	go func() {
		defer close(x2Chan)
		for x := range xChan {
			x2Chan <- x * 2
		}
	}()

	for x2 := range x2Chan {
		fmt.Println(x2)
	}
}
