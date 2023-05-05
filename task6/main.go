package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//мне кажется, что я в полной мере продемонстрировал возможные способы выхода из горутин в предыдущих заданиях (особенно в 5)
	ctx, ctxCancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Выход с помощью контекста")
		}
	}()
	fmt.Println("Нажмите enter, чтоб прервать горутину")
	fmt.Scanln()
	ctxCancel()

	quit := make(chan struct{})
	go func() {
		select {
		case <-quit:
			fmt.Println("Выход с помощью канала-оповещения")
		}
	}()
	fmt.Println("Нажмите enter, чтоб прервать горутину")
	fmt.Scanln()
	quit <- struct{}{}

	go func() {
		for {
			if _, ok := <-quit; !ok {
				fmt.Println("Выход с помощью закрытия канала")
				return
			}
		}
	}()
	fmt.Println("Нажмите enter, чтоб прервать горутину")
	fmt.Scanln()
	close(quit)

	go func() {
		select {
		case <-time.NewTimer(time.Second).C:
			fmt.Println("Выход с помощью таймеров, afterFunc, etc.")
		}
	}()

	time.Sleep(2 * time.Second)
}
