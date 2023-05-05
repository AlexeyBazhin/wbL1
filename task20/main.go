package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog       sun"
	str = reversePointers(str)
	fmt.Println(str)
	str = reverseBuffer(str)
	fmt.Println(str)
	//str = "snow    dog sun    "
	str = reverseSpaces(str)
	fmt.Println(str)
}

func reversePointers(str string) string {
	words := strings.Split(str, " ")
	//строка с несколькими пробелами подряд (snow   dog  sun) 3 и 2 пробела
	//запишется в words как {"snow", "", "", dog, "", sun}
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ") //позволит разделить элементы пробелами (за счет чего получается их изначальное количество)
}

func reverseBuffer(str string) string {
	words := strings.Split(str, " ")
	builder := new(strings.Builder)
	for i := len(words) - 1; i >= 0; i-- {
		builder.WriteString(words[i])
		if i > 0 {
			builder.WriteRune(' ')
		}
	}
	return builder.String()
}

func reverseSpaces(s string) string {
	builder := new(strings.Builder)
	i, j := len(s)-1, len(s)-1
	for i >= 0 {
		for i >= 0 && s[i] == ' ' {
			i--
			j--
			builder.WriteByte(' ') //если есть желание перевернуть пробелы)))
		}
		if i < 0 {
			break
		}
		for i >= 0 && s[i] != ' ' { //ищем пробел
			i--
		}
		builder.WriteString(s[i+1 : j+1]) //записываем слово, которое прошли перед пробелом
		j = i
	}
	return builder.String()
}
