package main

import (
	"fmt"
	"sync"
	"time"
)

// CONSTANTS
const hunger = 3

// VARIABLES

var philosophers = []string{"Plato", "Socrates", "Aristotle", "Pascal", "Locke"}
var orderOfFinish = make(chan string, len(philosophers))
var sleepTime = 1 * time.Second
var eatTime = 3 * time.Second
var thinkTime = 1 * time.Second
var wg sync.WaitGroup

func diningProblem(philosopher string, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()
	// DINING PROBLEM

	// Print a message
	fmt.Printf("%s is seated at the table \n", philosopher)
	time.Sleep(sleepTime)

	for i := hunger; i > 0; i-- {
		fmt.Printf("%s is hungry and want to eat \n", philosopher)
		time.Sleep(sleepTime)

		//  Lock both forks
		leftFork.Lock()
		fmt.Printf("\t%s has acquired the left fork \n", philosopher)
		rightFork.Lock()
		fmt.Printf("\t%s has acquired the right fork \n", philosopher)

		// Print a message when philosopher has both forks
		fmt.Printf("%s is eating \n", philosopher)
		time.Sleep(eatTime)

		fmt.Println(philosopher, "is thinking")
		time.Sleep(thinkTime)

		// Unlock the mutexes
		rightFork.Unlock()
		fmt.Printf("\t%s has released the right fork \n", philosopher)
		leftFork.Unlock()
		fmt.Printf("\t%s has released the left fork \n", philosopher)
		time.Sleep(sleepTime)
	}
	fmt.Printf("\t%s is satisfied \n", philosopher)
	time.Sleep(sleepTime)

	fmt.Printf("\t%s left the table \n", philosopher)
	orderOfFinish <- philosopher
}

func main() {
	// INTRO

	fmt.Println("The dining philosophers problem")
	fmt.Println("-------------------------------")
	fmt.Println()

	leftFork := &sync.Mutex{}

	wg.Add(len(philosophers))
	// SPAWN A GOROUTINE FOR EACH PHILOSOPHER
	// for i := 0; i < len(philosophers); i++ {
	// 	go diningProblem(philosophers[i])
	// }
	for _, philosopher := range philosophers {
		rightFork := &sync.Mutex{}
		go diningProblem(philosopher, leftFork, rightFork)

		leftFork = &sync.Mutex{}
	}

	wg.Wait()

	fmt.Println()
	fmt.Println("The philosophers are satisfied")
	fmt.Println()
	fmt.Println("The order of finish is:")
	for i := 0; i < len(philosophers); i++ {
		fmt.Println("\t", <-orderOfFinish)
	}

}
