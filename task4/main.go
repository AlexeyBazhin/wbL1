package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"sync"
	"syscall"
)

var (
	workerWg *sync.WaitGroup
	dataWg   *sync.WaitGroup
	writer   *bufio.Writer
)

func startWorker(workerNum int, workerChan chan int) {
	fmt.Printf("Воркер под номером %v начал работу\n", workerNum)
	for data := range workerChan { //каждый воркер слушает общий канал с данными		

		fmt.Printf("Воркер %v обработал операцию: %v\n", workerNum, data)
		//dataWg.Done()
		runtime.Gosched() //оповещаем планировщик задач о том, что нужно переключить контекст выполнения
		//если закомментировать эту строчку, то будет видно, что переключение между горутинами происходит редко
		//одна из них не будет освобождать ресурсы и давать другим горутинам поработать
	}
	fmt.Printf("Воркер под номером %v завершил работу\n", workerNum)
	workerWg.Done()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var workerCount int
	fmt.Println("Введите количество воркеров: ")
	fmt.Fscan(reader, &workerCount)

	writer = bufio.NewWriter(os.Stdout)
	workerChan := make(chan int, 3) //буфер можно задать любого желаемого размера
	workerWg = &sync.WaitGroup{}
	dataWg = &sync.WaitGroup{}
	for i := 1; i <= workerCount; i++ { //стартуем воркеров в отдельных горутинах
		workerWg.Add(1)
		go startWorker(i, workerChan)
	}

	//создаем канал-оповещение о выходе.
	//тип канала - пустая структура, т.к. она является самым легковестным типом в языке и сама по себе весит 0 байт
	quit := make(chan struct{})
	go func() {
		//создаем канал типа оповещение ОС,
		shutdown := make(chan os.Signal, 1)                    //он должен быть буферизированным, т.к. может возникнуть дедлок, когда система отправит какой-либо сигнал, но нигде в коде он уже приниматься из канала не будет
		signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM) //указываем какого сигналы какого типа нужно закидывать в канал
		for {
			select {
			case sig := <-shutdown: //в канал пришел сигнал
				writer.WriteString(strings.Join([]string{"\nПрограмма завершилась с сигналом:", sig.String(), "\n"}, " "))
				//fmt.Printf("\nПрограмма завершилась с сигналом: %s\n", sig)
				quit <- struct{}{} //наполняем канал-оповещение, чтоб другая горутина поняла, что нужно завершить свои операции
				return
			}
			//я оставил for select (хотя можно было бы использовать код ниже)
			//так как обычно в полноценных программах бывает кейс <-ctx.Done(), в случае возникновения которого нужно сделать немного другие действия
			// sig := <-shutdown
			// quit <- struct{}{}
			// fmt.Printf("\nПрограмма завершилась с сигналом: %s\n", sig)
		}
	}()

	data := 0
	for {
		select {
		case <-quit: //в данном случае понимаем, что пришел сигнал на выход из программы
			//dataWg.Wait()
			close(workerChan) //закрываем канал с данными, чтобы воркеры завершили свою работу
			workerWg.Wait()
			writer.Flush()
			return
		default:
			//dataWg.Add(1)
			workerChan <- data //записываем данные в канал
			data++
		}
	}

}
