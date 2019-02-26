package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var output, input int

	time.Sleep(2 * time.Second)
	fmt.Printf("GAMES START NOW\n")
	fmt.Printf("...\n")
	time.Sleep(2 * time.Second)
	fmt.Printf("...\n")
	fmt.Printf("NOW SELECT CORRECT NUMBER FOR INPUT\n")
	fmt.Printf("1.ROCK\n 2.PAPER\n 3.SCISSOR")
	fmt.Scanf("%d", input)

	output = rand.Intn(3)

	if output == 1 && input == 1 {
		fmt.Printf("You Choose=ROCK\n CPU Choose=ROCK")
		fmt.Printf("...")
		time.Sleep(2 * time.Second)
		fmt.Printf("...")
		fmt.Printf("You Choose=ROCK\n CPU Choose=ROCK")
		fmt.Printf("TIE try again")
	}
	if output == 1 && input == 2 {
		fmt.Printf("You Choose=ROCK\n CPU Choose=PAPER")
		fmt.Printf("...")
		time.Sleep(2 * time.Second)
		fmt.Printf("...")
		fmt.Printf("YOU WIN")

	}
	if output == 1 && input == 3 {

		fmt.Printf("YOU LOOSE")
		fmt.Printf("You Choose=ROCK\n CPU Choose=SCISSOR")
		fmt.Printf("...")
		time.Sleep(2 * time.Second)
		fmt.Printf("...")
		fmt.Printf("YOU LOOSE")
	}
	if output == 2 && input == 1 {
		fmt.Printf("You Choose=PAPER\n CPU Choose=ROCK")
		fmt.Printf("...")
		time.Sleep(2 * time.Second)
		fmt.Printf("...")
		fmt.Printf("YOU LOOSE")
	}
	if output == 2 && input == 2 {
		fmt.Printf("You Choose=PAPER\n CPU Choose=PAPER")
		fmt.Printf("...")
		time.Sleep(2 * time.Second)
		fmt.Printf("...")
		fmt.Printf("TIE try again")

	}
	if output == 2 && input == 3 {
		fmt.Printf("You Choose=PAPER\n CPU Choose=SCISSOR")
		fmt.Printf("...")
		time.Sleep(2 * time.Second)
		fmt.Printf("...")
		fmt.Printf("YOU win")
	}
	if output == 3 && input == 1 {
		fmt.Printf("You Choose=SCISSOR \n CPU Choose=PAPER")
		fmt.Printf("...")
		time.Sleep(2 * time.Second)
		fmt.Printf("...")
		fmt.Printf("YOU win")
	}
	if output == 3 && input == 2 {
		fmt.Printf("You Choose=SCISSOR\n CPU Choose=PAPER")
		fmt.Printf("...")
		time.Sleep(2 * time.Second)
		fmt.Printf("...")
		fmt.Printf("YOU LOOSE")
	}
	if output == 3 && input == 3 {
		fmt.Printf("You Choose=SCISSOR\n CPU Choose=SCISSOR")
		fmt.Printf("...")
		time.Sleep(2 * time.Second)
		fmt.Printf("...")
		fmt.Printf("TIE try again")
	}

}
