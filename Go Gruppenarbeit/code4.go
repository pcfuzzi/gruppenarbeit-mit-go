package main

import "fmt"

func main() {
	var start int
	var end int
	fmt.Print("Geben Sie den Startwert ein: ")
	fmt.Scan(&start)
	fmt.Print("Geben Sie den Endwert ein: ")
	fmt.Scan(&end)
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	fmt.Println("Die Summe aus dem Schleifendurchlauf ist:", sum)
}