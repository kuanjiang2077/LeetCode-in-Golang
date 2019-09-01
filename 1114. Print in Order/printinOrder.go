package main

import "fmt"

func main() {
	cs := make([]chan int, 3)
	for i := range cs {
		cs[i] = make(chan int)
	}
	go first(cs[0])
	go second(cs[1])
	go third(cs[2])

	for i := range cs {
		cs[i] <- 1
		<- cs[i]
	}
}

func first(c chan int) {
	<- c
	fmt.Println("one")
	c <- 1
}

func second(c chan int) {
	<- c
	fmt.Println("two")
	c <- 1
}

func third(c chan int) {
	<- c
	fmt.Println("three")
	c <- 1
}
