package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Print("Gib dein Alter ein: ")
	fmt.Scan(&num1)
	fmt.Print("Wie alt ist deine Mutter: ")
	fmt.Scan(&num2)

	sum := num1 + num2
	fmt.Println("Du bist", num1,"Jahre und deine Mutter", num2, "Jahre, zusammen seid ihr", sum, "Jahr alt")

	var value int
	fmt.Print("Bitte gib dein Alter an: ")
	fmt.Scan(&value)

	if value >= 18 {
		fmt.Println("Du darfst alles kaufen was du willst!")
	} else {
		fmt.Println("Du darfst nicht alles kaufen was du willst.")
	}
}

