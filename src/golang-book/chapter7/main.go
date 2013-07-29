package main

import "fmt"

func main(){
	fmt.Println(add(3,4))
	fmt.Println(multiReturn())
}

func add(a int, b int) int{
	return a + b
}

func multiReturn() (int, int){
	return 3, 5
}
