package main

import (
	"context"
	"fmt"
	"time"
)

// я уже подробно использовал и описал необходимые для реализации подходы в задании 5
func main() {
	mySleepTimer(1 * time.Second)
	fmt.Println("проснулись")
	mySleepTimeAfter(1 * time.Second)
	fmt.Println("снова проснулись")
	mySleepCtxTimeout(1 * time.Second)
	fmt.Println("еще раз проснулись")
}
func mySleepTimer(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}
func mySleepTimeAfter(duration time.Duration) {
	<-time.After(duration)
}
func mySleepCtxTimeout(duration time.Duration) {
	ctx, _ := context.WithTimeout(context.Background(), duration)
	<-ctx.Done()
}
