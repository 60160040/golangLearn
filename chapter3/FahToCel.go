package main

import "fmt"

func main() {
	var Fah int
	fmt.Scanf("%d", &Fah)
	c := float64((Fah - 32)) * 5 / 9 //int can't procass with float
	fmt.Println(c)
}
