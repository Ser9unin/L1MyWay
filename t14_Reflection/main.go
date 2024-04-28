package main

import (
	"fmt"
	"reflect"
)

// Вариант 1 Рефлексия
func typeReflection(v interface{}) []string {
	typ := "type: " + reflect.TypeOf(v).String()
	kind := ", kind: " + reflect.TypeOf(v).Kind().String()
	res := []string{typ, kind}
	return res //выводит всю правду матку
}

// Вариант 2 Sprintf
func typeSprintf(v interface{}) string {
	return fmt.Sprintf("%T", v) //выводит тип который виден на верхнем уровне, не проваливается в структуру
}

// Вариант 3 switch интерфейса по типам
func typeSwitch(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan string:
		return "chan string"
	default:
		return "unknown"
	}
}

type Mytype struct {
	val int
}

func main() {
	b := true
	strType := "data"
	intType := 1
	channel := make(chan string)
	mytype := Mytype{
		val: 2,
	}
	//Вариант 1 Рефлексия
	fmt.Println("Вариант 1 Рефлексия:")
	fmt.Println(typeReflection(b))
	fmt.Println(typeReflection(strType)) //print string
	fmt.Println(typeReflection(intType)) //print int
	fmt.Println(typeReflection(channel))
	fmt.Println(typeReflection(mytype))

	//Вариант 2 Sprinf внутри так же рефлексия
	fmt.Println("Вариант 2 Sprinf внутри так же рефлексия:")
	fmt.Println(typeSprintf(b))
	fmt.Println(typeSprintf(strType))
	fmt.Println(typeSprintf(intType))
	fmt.Println(typeSprintf(channel))
	fmt.Println(typeSprintf(mytype))

	//Version 3
	fmt.Println("Вариант 3 switch интерфейса по типам:")
	fmt.Println(typeSwitch(b))
	fmt.Println(typeSwitch(strType))
	fmt.Println(typeSwitch(intType))
	fmt.Println(typeSwitch(channel))
	fmt.Println(typeSwitch(mytype))
}
