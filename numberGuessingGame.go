package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//This code creates random number and place as variable
	rand.Seed(time.Now().UnixNano())
	hiddenNumber := int64(rand.Intn(100))
	//Insert code here
	var userNum int64
	fmt.Println("Guess the number between 1 to 100!")

	for i := 5; i >= 0; i-- {
		if i > 0 {
			fmt.Printf("You have %v tries!\n", i)
			fmt.Println("Enter 101 to exit program.")
			fmt.Println("Enter number: ")
			fmt.Scanln(&userNum)

			if userNum == hiddenNumber {
				fmt.Println("You guessed correct!")
				break
			} else if userNum != hiddenNumber {
				fmt.Println("You guessed wrong!")
			} else if userNum == 101 {
				fmt.Println("Goodbye!")
				break
			} else {
				fmt.Print("Error. That is not a number!")
			}
		} else {
			fmt.Println("You ran out of tries! You Lose!")
		}
	}
	fmt.Println("Exiting program.")

}
