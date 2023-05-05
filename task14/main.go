package main

import (
	"fmt"
	"reflect"
)

func main() {
	typeAssertion(5)
	typeAssertion("строка")
	typeAssertion(true)
	ch := make(chan int, 1)
	ch <- 23
	typeAssertion(ch)
	typeAssertion(struct{}{})

	// withReflect(&struct{
	// 	G int
	// 	S string
	// }{5, "fdd"})

	withReflect(5)
	withReflect("stroka")
	withReflect(true)
	withReflect(ch)
}

func typeAssertion(i interface{}) {
	switch value := i.(type) {
	case int:
		fmt.Printf("Тип:%T;Значение в 2 СС: %b\n", value, value)
	case string:
		fmt.Printf("Тип:%T;Значение: %q\n", value, value)
	case bool:
		fmt.Printf("Тип:%T;Значение: %t\n", value, value)
	case chan int:
		fmt.Printf("Тип:%T;Значение: %v\n", value, <-value)
	default:
		fmt.Println("Не знаю")
	}
}

func withReflect(i interface{}) {
	// val := reflect.ValueOf(i)
	// fmt.Println(val.Type())
	// fmt.Println(val.Kind())
	// fmt.Println(val.Type().Kind())
	
	fmt.Printf("Type:%v;Kind:%v Value:%v\n", reflect.TypeOf(i),reflect.ValueOf(i).Kind(), reflect.ValueOf(i))
}
