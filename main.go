package main

import "fmt"

func addOne(a int) {
	a++
	fmt.Println("Внутри функции: ", a)
}

func addOnePtr(p *int) int {
	result := (*p) + 1
	return result
}

func main() {
	x := 42

	p := &x

	fmt.Printf("x = %d\t p = %v\n", x, p)
	fmt.Printf("Значение по адресу, на который указывает p: %d\n", *p)
	*p++
	fmt.Printf("Значение по адресу, на который указывает p: %d\n", *p)

	fmt.Println("--- Указатели в функциях ---")
	addOne(x)
	fmt.Println("После работы функции: ", x)
	fmt.Println(addOnePtr(&x))
}
