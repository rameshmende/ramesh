package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Switch cases in golang")

	rand.Seed(time.Now().UnixNano())
	DiceNumber := rand.Intn(6)+1
	switch DiceNumber {
	case 1 :
		fmt.Println("the value dice is 1 open ")

	case 2 :
		fmt.Println("the role of dice is 2 to move 2 spot")	

	case 3 :
		fmt.Println("the value of dice role is 3 to move 3 spot")	
        fallthrough
	case 4 :
		fmt.Println("the value of dice is 4 to move 4 spot")	
        fallthrough
	case 5 :
		fmt.Println("value of dice is 5 to move 5 spot")	

	case 6 :
		fmt.Println("value of dice is 6 to move 6 spot and again role")
		
	default :
	fmt.Println("What was that!")	
	}
}