package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	
	scanner := bufio.NewScanner(os.Stdin)

	high := 100
	low := 1
	attempts := 0
	lieDetector := 0

	fmt.Println("Think about number between", low, "and", high)
	fmt.Println("Press Enter to start")
	scanner.Scan()



	for {
		guess := (high + low) / 2
		attempts++
		
		if lieDetector == guess{
			fmt.Println("You are cheating or made a mistake")
			break
		}

		lieDetector = guess
		
		fmt.Println("Now i think the number is ", guess)
		fmt.Println("(a) Number is lower")
		fmt.Println("(b) Number is higher")
		fmt.Println("(c) Number is correct")

		scanner.Scan()

		r := scanner.Text()

		if r == "a"{
			high = guess -1 
		}else if r == "b"{
			low = guess + 1
		}else if r == "c"{
			fmt.Println("I won!")
			break
		}else{
			fmt.Println("Invalid choice")
		}

		
	} 
	fmt.Println("-------------------------------------")
	fmt.Print("Attempts: ",attempts)
}