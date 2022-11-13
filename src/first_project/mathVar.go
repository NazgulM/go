package main 

import "fmt"

func main() {
	number := 10
	number++
	number +=1
	number = number + 1
	const s string = "constant"
	fmt.Println(rune(number))
	fmt.Println(20*3)
	fmt.Println(55.21/5)
	fmt.Println(9%2) //modulus returns a remainder
	fmt.Println(10%5)
	fmt.Println(s)
}