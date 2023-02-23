package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a value: ")
	input, _ := reader.ReadString('\n')
	value, _ := strconv.Atoi(input[:len(input)-1])

	if value >= 18 {
		fmt.Println("Du darfst alles kaufen, was du willst!")
	} else if value >= 7 && value < 18 {
		fmt.Println("Du darfst nur im Rahmen deines Taschengeldes einkaufen und mit Erlaubnis deiner Eltern!")
	} else if value > 0 && value < 7 {
		fmt.Println("Du darfst leider nichts kaufen, was dir nicht deine Eltern besorgen!")
	} else {
		fmt.Println("Invalid value entered.")
	}
}
