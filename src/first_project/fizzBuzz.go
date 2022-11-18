package main

import "fmt"

//Fizz Buzz print numbers 0-100, but if 
//1. The number is divisible by 3, print "Fizz"
//2. If the number is divisible by 5, print "buzz"
//3. If the number is divisible by 3 and 5, print "fizzBuzz"

func main() {
	for num :=0; num <100; num++ {
		if num%3==0 && num%5==0 {
			fmt.Println("FizzBuzz")
		} else if num%3==0 {
			fmt.Println("Fizz")
		} else if num%5==0  {
			fmt.Println("Buzz")
		}else {
			fmt.Println(num)
		}
		
	}
}