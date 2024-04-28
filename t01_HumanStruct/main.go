package main

import "fmt"

type Human struct {
	Name       string
	SecondName string
	Age        int
}

// функтции описанные для Human технически можно было записать так
// func SomeFunc (h *Human, d data){} запись функций в формате func (p *pointer) FuncName(args...)
// сделана для удобства восприятия разработчиков.
func (h *Human) SetName(name string) {
	h.Name = name
}

func (h *Human) SetSecondName(sname string) {
	h.SecondName = sname
}

func (h *Human) SetAge(age int) {
	h.Age = age
}

func (h *Human) GetName() string {
	return h.Name
}

func (h *Human) GetSecondName() string {
	return h.SecondName
}

func (h *Human) GetAge() int {
	return h.Age
}

func (a *Action) SetName(name string) {
	a.Name = name
}

type Action struct {
	Name string
	Human
	Canwork   bool
	SickLeave bool
}

func main() {
	action := &Action{
		Name: "Vasya",
		Human: Human{
			Name:       "Petya",
			SecondName: "Petrov",
			Age:        35,
		},
	}

	// Обращение к полям структуры
	// при наличии одинаковых полей приоритет у поля полее высокого уровня,
	//в этом случае что бы получить данные из вложенной структуре нужно к ней обратиться явно
	fmt.Println("Имя на верхнем уровне", action.Name,
		" Имя во вложенной структуре", action.Human.Name)
	// если у поля в верхнеуровневой структуре нет дублера, то обращение будет к полю во вложенной структуре
	fmt.Println("Возраст транслируется на верхний уровень", action.Age)
	// запись выше будет аналогична такой
	fmt.Println("Получили возраст явно ", action.Human.Age)

	// Вызов методов

	// методы вложенной структуры наследуются родительской
	action.SetAge(55)
	fmt.Println("Задали возраст через функцию вложенной структуры", action.Age)

	// при наличии одинаковых методов приоритет у метода структуры верхнего уровня
	action.SetName("UltraLord3000")
	fmt.Println("Имя на верхнем уровне", action.Name,
		" Имя во вложенной структуре", action.Human.Name)

	// аналогичный вызов в отношении вложенной структуры нужно делать явно
	action.Human.SetName("Peasant")
	fmt.Println("Имя на верхнем уровне", action.Name,
		" Имя во вложенной структуре", action.Human.Name)
}
