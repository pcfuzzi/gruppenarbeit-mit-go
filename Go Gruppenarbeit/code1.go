package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("Geben Sie den ersten Wert ein: ")
	fmt.Scan(&a)
	fmt.Print("Geben Sie den zweiten Wert ein: ")
	fmt.Scan(&b)
	var c int = a + b
	fmt.Println("Das Ergebnis der Addition ist:", c)
	
}
