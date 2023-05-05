package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	getSquaresChan(arr)
	wg := &sync.WaitGroup{}
	getSquareFast(arr, wg)
	printSquareWg(arr, wg)
	getSquareChanWg(arr, wg)

}

// Получение квадратов с использованием wg и канала по мере их поступления. Как по мне, самый эффективный вариант из предложенных.
func getSquareFast(arr []int, wg *sync.WaitGroup) {
	fmt.Println("\nWg + any chan:")

	anyChan := make(chan int) //(может быть не буфер., размер может не совпадать с размером массива)
	for _, elem := range arr {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			anyChan <- number * number
		}(elem)
	}

	//Главная фишка - ожидание WG, после чего закрытие канала в отдельной горутине.
	go func() {
		fmt.Println("Waiting") //Данная строчка может напечататься после печати каких-либо квадратов (ниже по коду).
		wg.Wait()
		close(anyChan) //Закрываем канал, чтобы можно было использовать range и не зависить по размеру от других программных сущностей.
	}()

	fmt.Println("Check")
	for square := range anyChan {
		//Здесь мы обрабатываем данные по мере их поступления.
		//Не блокируемся на ожидание перед этим.
		//И можем отдавать данные порционно если укажем у канала anyChan размер буфера.
		fmt.Println(square)
	}
}

// Печать данных внутри горутины с использованием только WG
func printSquareWg(arr []int, wg *sync.WaitGroup) {
	fmt.Println("\nWg:")

	for _, elem := range arr {
		wg.Add(1)
		go func(number int, wg *sync.WaitGroup) {
			defer wg.Done()              //Откладываем декремент счетчика waitGroup до окончания выполнения функции. Т.е. когда данные будут напечатаны
			fmt.Println(number * number) //Минус примера - оперируем данными внутри горутины, а не передаем их в родительскую.
		}(elem, wg)
	}
	wg.Wait() //Блокируемся, пока все не напечатаются
}

// Получение канала с использование waitGroup и буф. канала
func getSquareChanWg(arr []int, wg *sync.WaitGroup) {
	fmt.Println("\nWg + buf chan:")

	bufChan := make(chan int, len(arr)) //Здесь буфер строгого размера
	for _, elem := range arr {
		wg.Add(1)
		//каналы могут быть однонаправленными - в данном случае только на запись внутри анонимной функции.
		//конечно, можно было бы использовать замыкание, использовать внутри bufChan, но я хотел показать однонаправленность
		go func(number int, squareChan chan<- int) {
			defer wg.Done() //Откладываем декремент счетчика waitGroup до окончания выполнения функции. Т.е. когда данные уже будут записаны в канал.
			squareChan <- number * number
		}(elem, bufChan)
	}
	wg.Wait()      //блокируемся до момента, когда счетчик wg станет 0, т.е. когда все горутины запишут свое значение в канал.
	close(bufChan) //Закрываем канал, запись в него блокируется.
	// В этом и заключается минус данного подхода. Мы не можем отправлять данные другими "порциями" - только полностью заполнить буфер.
	for square := range bufChan {
		fmt.Println(square)
	}
}

// Получение квадратов с использованием только канала
func getSquaresChan(arr []int) {
	fmt.Println("Chan:")

	ch := make(chan int) //канал в данном примере может быть как буферизированный, так и не буферизированный
	for _, elem := range arr {
		go func(number int) {
			ch <- number * number //отправляем данные в канал
		}(elem)
	}

	//Минус данного примера заключается в том, что мы строго зависим от длины массива.
	//Если бы квадраты рассчитывались не у всех элементов, а по какому-либо условию, то пришлось бы заводить доп переменную-счетчик.
	//И тут применять ее вместо len(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(<-ch) //получаем данные из канала
	}
}
