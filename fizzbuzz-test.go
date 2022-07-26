package main

import "fmt"

func main() {
	for count := 1; count <= 100; count++ {
		if count%3 == 0 && count%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			if count%3 == 0 {
				fmt.Println("Fizz")
			} else if count%5 == 0 {
				fmt.Println("Buzz")
			}
		}
	}
}
