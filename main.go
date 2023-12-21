package main

import (
	"example/alGOrithm/linkedlist"
	"fmt"
)

func main() {
	mylist := linkedlist.LinkedList{}
	mylist.Append(2)
	mylist.Append(1)
	mylist.Append(3)
	mylist.Prepend(9)
	mylist.Prepend(4)
	mylist.PrintList()
	len := mylist.GetLength()
	fmt.Println(len)
}
