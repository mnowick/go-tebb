package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	ex "github.com/mnowick/go-tebb/examples"
)

func main() {
	fmt.Println("Greetings, Unifiers!")
	fmt.Println("Which option would you like to select?")
	fmt.Println("-----------------------------------------")
	fmt.Println("1. Hello, World!")
	fmt.Println("2. Error Handling!")
	fmt.Println("3. Concurrency!")
	fmt.Println("4. Defer!")
	fmt.Println("0. Exit!")
	fmt.Println("-----------------------------------------")

	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text == "0" {
			fmt.Println("Hope you had fun!")
			break
		}
		switch text {
		case "1":
			ex.HelloUnify()
		case "2":
			var f float64
			fmt.Println("Enter a number")
			fmt.Scanf("%f", &f)
			ex.Error(f)
		case "3":
			ex.ConcurrencyMenu()
		case "4":
			ex.DeferEx()
		default:
			fmt.Println("Not a recognized input.")
		}
		fmt.Println("")
		fmt.Println("Would you like to pick another option?")
	}
}
