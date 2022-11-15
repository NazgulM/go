package main

import "fmt"

func main() {
	fmt.Println("5==5 is", 5==5)
	fmt.Println("5!=6 is", 5!=6)
	fmt.Println("2000 > 5 is", 2000 >5)
	fmt.Println("2000<5 is ", 2000 < 5)
	fmt.Println("5>=5", 5>=5)
	fmt.Println("Strong = strong", "strong" == "Strong")
	//Check if 5 is > num OR if num ==10
	num :=13
	fmt.Println("5> num(13) or num(13) ==10 ",  5 > num || num ==10)
}