// lvcl.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var output int
	var m, f string
	fmt.Printf("Enter Male Name\n")
	fmt.Scanln(&m)
	fmt.Printf("Enter Female Name\n")
	fmt.Scanln(&f)
	fmt.Printf("Game is Starting\n")
	fmt.Printf(".")
	time.Sleep(1 * time.Second)
	fmt.Printf(".")
	time.Sleep(1 * time.Second)
	fmt.Printf(".")
	time.Sleep(1 * time.Second)
	fmt.Printf(">")
	time.Sleep(1 * time.Second)
	fmt.Printf("-")
	time.Sleep(1 * time.Second)
	output = rand.Intn(100)
	fmt.Printf("LOVE IN BETWEEN THEM = %d PERCENT", output)
}
