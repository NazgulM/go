package main

import "fmt"

func main() {
	for num :=0; num <= 100; num++ {
		fmt.Println(num)
	}

	for num2 :=1; num2<10; num2++{
		fmt.Println(num2)
	}

	names := []string{"James", "Aaron", "Tom", "Brady"}
	for x:=0; x < len(names); x++ {
		fmt.Println(names[x])
	}
	names2 := []string {"Naza", "Chika", "Aruuke", "Nursultan", "Bakai", "Naza"}	
	for x2:= 0; x2 <len(names2); x2++ {
		if names2[x2] =="Naza" {
			fmt.Println("Hello Naza")
	}
}
}