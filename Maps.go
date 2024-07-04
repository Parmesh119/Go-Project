package main

import "fmt"

func Maps() {
	// creating maps in go
	mp := make(map[string]int)
	mp["first"] = 1
	fmt.Print(mp, "\n")
}