package main

import (
	"fmt"
	"unicode"
)

// также можно было бы пройтись циклом в цикле по рунам, сравнивая каждую со следующими.
// Но эффективное решение O(n) - с использованием мапы
func main() {
	fmt.Println(allUnique("abcd"))
	fmt.Println(allUnique("abCdefA"))
	fmt.Println(allUnique("aabcd"))
	fmt.Println(allUnique("Абв"))
	fmt.Println(allUnique("АбвБг"))
}

// использую мапу, т.к. по ключу можно легко понять, встречался ли элемент
// здесь нам нужен только ключ, поэтому значения мапы - пустые структуры, которые весят 0 байт
func allUnique(str string) bool {
	seen := make(map[rune]struct{})
	for _, strRune := range str {
		strRune = unicode.ToLower(strRune)
		//strRune = toLowerRune(strRune) //собственная реализация
		if _, ok := seen[strRune]; ok {
			return false
		}
		seen[strRune] = struct{}{}
	}
	return true
}

// также вместо пакета unicode, которые предоставляет функции для всех своих символов
// можно было бы использовать свою реализацию, делая следующий вызов: strRune = rune(toLowerByte(byte(strRune)))
// например, для латиницы:
func toLowerByte(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b - 'A' + 'a'
	}
	return b
}

// но если делать реализации и для других языков, чьи символы весят >1 байта
// то пришлось бы делать отдельные кейсы уже для рун
func toLowerRune(r rune) rune {
	switch {
	case 'A' <= r && r <= 'Z':
		r += 'a' - 'A' //разница между строчной и заглавной буквой в ascii одинакова (32)
		return r
	case r >= 'А' && r <= 'Я':
		r += 'а' - 'А'
	}
	//и тд
	return r
}
