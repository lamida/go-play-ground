package main

import "fmt"

func main(){
	fmt.Println("how many number to input?")
	var input int
	fmt.Scanf("%d", &input)
	x := make([]int, input)
	fmt.Println(x)
	for i := 0; i < input; i++ {
		fmt.Println("enter number", i + 1)
		fmt.Scanf("%d", &x[i])
	}
	fmt.Println(x)
}
