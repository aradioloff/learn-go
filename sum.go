package main

import (
	"fmt"
)

func main(){
	var a, b int
	fmt.Println("Imput A")
	fmt.Scan(&a)
	fmt.Println("Imput B")
	fmt.Scan(&b)
	sum := a+b
	fmt.Println("Sum = ", sum)
}
