package main

import (
	"fmt"
	"time"
	"sync"
)

type s struct {
	firstName string
	lastName string
	age uint
}

var wg = sync.WaitGroup{}

func main() {
	var data = s {
		firstName: "Parmesh",
		lastName: "Bhatt",
		age: 20,
	}
	fmt.Print(data)
	fmt.Print("The name of the person is", data.firstName, data.lastName, ". The age of the person is ", data.age, "\n")
	wg.Add(1)
	go timeOut()
	wg.Wait()
}

func timeOut() {
	time.Sleep(1 * time.Second)
	fmt.Print("abcd\n")
	wg.Done()
}