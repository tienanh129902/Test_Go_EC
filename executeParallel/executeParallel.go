package main

import "fmt"

func executeParallel(ch chan<- int, functions ...func() int) {
	for _, v := range functions {
		ch <- v()
	}
	close(ch)
	return
}
func exampleFunction(counter int) int {
	sum := 0
	for i := 0; i < counter; i++ {
		sum += 1
	}
	return sum
}
func main() {
	expensiveFunction := func() int { return exampleFunction(200000000) }
	cheapFunction := func() int { return exampleFunction(10000000) }
	ch := make(chan int)
	go executeParallel(ch, expensiveFunction, cheapFunction)
	for result := range ch {
		fmt.Printf("Result: %d\n", result)
	}
}
