package main

import "fmt"

type Set map[string]int

func (s Set) Add(item string) {
	s[item] = 1 //инициализация
}

//изменять объект внутри метода проверки будет неверно с логической точки зрения - она должна только возвращать результат проверки
//но для задачи я оставил так
func (s Set) Has(item string) bool {
	_, ok := s[item]
	if ok {
		s[item]++ //инкремент
	}
	return ok
}

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(Set)
	for _, word := range words {
		if !set.Has(word) { //в отличие от предыдущей задачи я делаю проверку есть ли элемент с таким ключом в множестве
			set.Add(word)
		}
	}

	for key, val := range set {
		fmt.Printf("%v: %v \n", key, val)
	}
}
