package main

import "fmt"

func main() {
	var age int
	fmt.Print("Geben Sie Ihr Alter ein: ")
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("Du darfst alles kaufen, was du willst!")
	} else if age >= 7 {
		fmt.Println("Du darfst nur im Rahmen deines Taschengeldes einkaufen und mit Erlaubnis deiner Eltern!")
	} else {
		fmt.Println("Du darfst leider nichts kaufen, was dir nicht deine Eltern besorgen!")
	}
}