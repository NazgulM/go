package main

import "fmt"

func main() {
	age := 29
	//Check if person is old enough to drink
	if(age >= 21) {
		fmt.Println("You may enter the bar, because you can drink")
	}
	age2 := 20
	if(age2 <= 21) {
		fmt.Println("You cannot enter the bar, because you are younger than 21")
	}
    
}