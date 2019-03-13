// ysk
package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	var i,j,a,b int
	stick:=21
	fmt.Println("Game is Starting")
	time.Sleep(2 * time.Second)
	F:
    for i=21;i<=0;i-a{
    	fmt.Print("Pickup the matchstick")
		
    	fmt.Scanln(&a)
		if a<4{
			fmt.Print("Wrong option choose under 4")
		}
		if i==0{
			fmt.Print("YOU LOOSE THE MATCH")
			exit()
		}
		
		for j=i;;j-b{
		fmt.Print("Now CPU pick MatchStick")
		b=rand.Intn(4)
		if j==0{
		fmt.Print("CPU Loose")
		}
		else{
		goto F
		}
		
    }
    }
}
