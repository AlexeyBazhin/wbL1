package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

var (
	quitHandler chan struct{}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var sec int
	fmt.Println("Введите N секунд: ")
	fmt.Fscan(reader, &sec)

	quitHandler = make(chan struct{})

	//6 вариантов:

	handler(sec, withTimer(sec))
	//handler(sec, withTimeAfter(sec))
	//handler(sec, withTimeAfterFunc(sec))
	//handler(sec, withTicker(sec))
	//handler(sec, withTick(sec))
	//handler(sec, withContextTimeout(sec))

}

// общая функция для всех вариантов решения задачи
func handler(sec int, dataChan <-chan int) {
	for {
		select {
		case <-quitHandler: // ожидает сигнала на завершение (получение данных вхолостую из quit и печать полученных данных)
			fmt.Println("Выполнено")
			return
		default:
			if x, ok := <-dataChan; ok { //необходимо делать проверку, т.к. горутина-отправитель может уже перестать высылать данные, а здесь мы их пытаемся принять
				fmt.Println(x)
			}
		}

	}
}

func withTimer(sec int) <-chan int {
	timer := time.NewTimer(time.Duration(sec) * time.Second) //структура таймер содержит внутри себя канал, в который запишутся данные после указаного промежутка времени
	dataChan := make(chan int, 1)
	go func() {
		counter := 0
		for {
			select {
			case <-timer.C: //получение данных из данного канала сигнализирует о том, что указанное время прошло
				close(dataChan) //закрываем, чтобы горутина-получатель не заблокировалась
				quitHandler <- struct{}{}
				return
			default:
				dataChan <- counter
				counter++
			}
		}
	}()
	return dataChan
}
func withTimeAfter(sec int) <-chan int {
	dataChan := make(chan int)
	go func() {
		timeAfterChan := time.After(time.Duration(sec) * time.Second) //возвращает сразу канал, без структуры
		counter := 0
		for {
			select {
			case <-timeAfterChan: // Поэтому мы не можем остановить остановить вручную. Пока данный таймер под капотом не выполнится, он не соберется сборщиком мусора.
				close(dataChan)
				quitHandler <- struct{}{}
				return
			default:
				dataChan <- counter
				counter++
			}
		}
	}()
	return dataChan
}
func withTimeAfterFunc(sec int) <-chan int {
	dataChan := make(chan int)
	timer := time.AfterFunc(time.Duration(sec)*time.Second,
		func() {
			quitHandler <- struct{}{}
		}) //то же самое, что и обычный таймер, только можем вызвать функцию, которая выполнится в отдельной горутине
	go func() {
		counter := 0
		for {
			select {
			case <-timer.C:
				close(dataChan)
				return
			default:
				dataChan <- counter
				counter++
			}
		}
	}()
	return dataChan
}

// код часто дублировался - вынес в отдельную функцию
func dataSenderStruct(dataChan chan int, quit <-chan struct{}) {
	counter := 0
	for {
		select {
		case <-quit:
			close(dataChan)
			quitHandler <- struct{}{}
			return
		default:
			dataChan <- counter
			counter++
		}
	}
}

func withTicker(sec int) <-chan int {
	dataChan := make(chan int)
	ticker := time.NewTicker(time.Second)
	quitTicker := make(chan struct{})
	//в отдельной горутине слушаем тикер и, после того как он тикнет sec раз, подаем сигнал на выход
	//но с тиками будет проблема, т.к. если код внутри range будет выполняться дольше секунды, то и таймаут случится не через нужное количество секунд
	go func() {
		secondCounter := 0
		for {
			<-ticker.C
			secondCounter++
			if secondCounter >= sec {
				quitTicker <- struct{}{}
				ticker.Stop() //останавливаем, иначе произойдет утечка
				return
			}
		}
	}()

	go dataSenderStruct(dataChan, quitTicker)

	return dataChan
}

func withTick(sec int) <-chan int {
	dataChan := make(chan int)
	tick := time.Tick(time.Second) //тик нужно использовать, если он должен работать вечно, потому что его нельзя остановить
	quitTicker := make(chan struct{})
	go func() {
		secondCounter := 0
		for {
			<-tick
			secondCounter++
			if secondCounter >= sec {
				quitTicker <- struct{}{}
				return
			}
		}
	}()

	go dataSenderStruct(dataChan, quitTicker)

	return dataChan
}

// отдельной реализации с WithDeadline не будет, т.к. это то же самое (тем более WithTimeout его сразу же использует под капотом)
func withContextTimeout(sec int) <-chan int {
	dataChan := make(chan int)
	ctxTimeout, _ := context.WithTimeout(context.Background(), time.Duration(sec)*time.Second)

	go dataSenderStruct(dataChan, ctxTimeout.Done())

	return dataChan
}
