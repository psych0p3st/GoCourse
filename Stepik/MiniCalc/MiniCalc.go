package main

import "fmt"

func main() {

	var a, b, res float64
	fmt.Scan(&a, &b)
	res = a + b
	fmt.Println("Сложение:", res)
	res = a - b
	fmt.Println("Вычитание:", res)
	res = a * b
	fmt.Println("Умножение:", res)
	res = a / b
	fmt.Println("Деление:", res)

}
