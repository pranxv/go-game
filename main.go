package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	secret := rand.Intn(100) + 1
	fmt.Println(secret)
	count := 5
	current := 1
	for current < count {
		fmt.Println("\n")
		fmt.Println("Enter your guess, the number should be between 1 and 100")
		var guess int
		fmt.Scan(&guess)
		fmt.Println(guess)
		if guess != secret {
			if int(math.Abs(float64(secret-guess))) > 10 {
				fmt.Printf("Your guess: %d is too far, you have %d tries left", guess, count-current)
			} else {
				fmt.Printf("Your guess: %d is close, take try between this range only, you have %d tries left", guess, count-current)

			}
		} else {
			fmt.Printf("YOU GOT IT RIGHT, in the %d try", current)
			break
		}
		current += 1
	}
}
