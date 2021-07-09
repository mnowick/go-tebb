package examples

import "fmt"

func DeferEx() {
	foo()
	fmt.Println("Returned normally from foo.")
}

func foo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in foo", r)
		}
	}()
	fmt.Println("Calling bar.")
	bar(0)
	fmt.Println("Returned normally from bar.")
}

func bar(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in bar", i)
	fmt.Println("Printing in bar", i)
	i++
	bar(i)
}
