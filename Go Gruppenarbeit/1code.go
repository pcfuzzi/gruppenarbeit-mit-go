package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the first number: ")
	num1Str, _ := reader.ReadString('\n')
	num1, _ := strconv.Atoi(num1Str[:len(num1Str)-1])

	fmt.Print("Enter the second number: ")
	num2Str, _ := reader.ReadString('\n')
	num2, _ := strconv.Atoi(num2Str[:len(num2Str)-1])

	sum := num1 + num2
	fmt.Println("The sum of the two numbers is:", sum)
}
