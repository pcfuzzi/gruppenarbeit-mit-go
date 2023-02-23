package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the start value: ")
	startStr, _ := reader.ReadString('\n')
	start, _ := strconv.Atoi(startStr[:len(startStr)-1])

	fmt.Print("Enter the end value: ")
	endStr, _ := reader.ReadString('\n')
	end, _ := strconv.Atoi(endStr[:len(endStr)-1])

	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	fmt.Println("The sum of the values from", start, "to", end, "is:", sum)
}
