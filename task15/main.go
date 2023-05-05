package main

import (
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	//в данном случае глобальная переменная будет ссылаться на подстроку большой строки
	//сама же подстрока ссылается на оригинальную
	//следовательно, большая строка не соберется сборщиком мусора и будет храниться в памяти
	//пока глобальная переменная ссылается на нее
	justString = v[:10]
}
func newSomeFunc() {
	hugeString := createHugeString(1 << 10)

	builder := new(strings.Builder)
	builder.WriteString(hugeString[:10])
	justString = builder.String()
}

func main() {
	someFunc()
	newSomeFunc()
	//таким образом, в новой функции создастся копия подстроки большой строки
	//после чего глобальная переменная будет ссылаться на копию
	//а большая строка соберется сборщиком мусора, т.к. на нее никто не ссылается
}

func createHugeString(strLen int) string {
	return "very big string"
}
