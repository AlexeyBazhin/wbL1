package main

import (
	"fmt"
)

type (
	Human struct {
		Name string //имя
		Age  int    //возраст
	}
	Action struct {
		*Human            //указатель на структуру Human. Был выбран указатель, так как с их помощью можно изменять значения у самих экземпляров структуры, а не их копий.
		stringerCount int //просто случайное поле для примера). В моем случае считает количество вызовов метода.
	}
)

// Да, внутри Action можно было оставить Human, а метод сделать на указатель, тогда бы преобразование к *Human при вызове метода происходило неявно. Но так как мы зачастую имеем дело с интерфейсами, то Human не реализовывал интерфейс с методом для указателя (или наоборот).
func (human *Human) AddAge(years int) {
	human.Age += years
}

// Имплементация интерфейса Stringer
func (human *Human) String() string {
	return fmt.Sprintf("%s is %v years old.", human.Name, human.Age)
}

// Так как я не до конца понял задание, я реализовал 2 своих видения. 
//В данном примере мы создаем метод, у которого должен быть такой же принцип работы, как и у метода встроенной структуры.
//Но есть дополнительный нюанс в бизнес логике.
func (action *Action) String() string {
	action.stringerCount++
	return fmt.Sprintf("\"%s\" (Данный метод вызвался в %v раз).", action.Human.String(), action.stringerCount)
}

func main() {
	//инициализируем, stringerCount - дефолтное значение 0
	action := &Action{
		Human: &Human{
			Name: "Alex",
			Age:  21,
		},
	}
	//Второе мое понимание - метод не реализован у самой структуры, но его можно вызывать напрямую, так как он есть у встроенной структуры.
	action.AddAge(1)    //для экземпляра Action вызывается метод AddAge встроенного *Human
	fmt.Println(action) //т.к. реализован интерфейс Stringer выведется результат метода String()
}
