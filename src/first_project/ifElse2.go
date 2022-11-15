package main

import "fmt"

func main() {
	age:= 20

	if age > 21 {
		fmt.Println("You may enter because you can drink")
	} else if age == 21 {
		fmt.Println("BlackJack! You may enter the bar")
	} else {
		fmt.Println("You shall not pass")
	}
}