package main

import (
	"example/alGOrithm/linkedlist"
	"fmt"
)

func main() {
	mylist := linkedlist.LinkedList{}

	len := mylist.GetLength()

	fmt.Println(len)
}
