package examples

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func sayExample() {
	go say("world")
	say("hello")
}

func unbufferedChannels() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func workerSimulation(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func waitGroups() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go workerSimulation(i, &wg)
	}

	wg.Wait()
}

func ConcurrencyMenu() {
	fmt.Println("Which example would you like to see?")
	fmt.Println("-----------------------------------------")
	fmt.Println("1. Basic Go Routine")
	fmt.Println("2. Unbuffered Channels")
	fmt.Println("3. Wait Groups")
	fmt.Println("0. Return to main menu")

	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "0" {
			fmt.Println("Returning to main menu")
			break
		}

		switch text {
		case "1":
			sayExample()
		case "2":
			unbufferedChannels()
		case "3":
			waitGroups()
		}
		fmt.Println()
		fmt.Println("Would you like to pick another option?")
	}
}
