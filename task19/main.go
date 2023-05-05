package main

import (
	"bytes"
	"fmt"
)

func main() {
	str := "главрыба"
	str = reversePointers(str)
	fmt.Println(str)
	str = reverseBuffer(str)
	fmt.Println(str)
}
func reversePointers(str string) string {
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 { //i++ j--
		r[i], r[j] = r[j], r[i] //два указателя, сдвигаются к центру, меняя соответствующие символы
	}
	return string(r)
}
func reverseBuffer(str string) string {
	r := []rune(str)
	buff := new(bytes.Buffer)
	for i := len(r) - 1; i >= 0; i-- {
		buff.WriteRune(r[i])
	}
	return buff.String()
}
