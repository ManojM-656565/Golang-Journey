package basics

import "fmt"

func Arrays() {
	// fmt.Print("Heoolp")
	// var arr [size]elementtype
	var numbers[5]int
	fmt.Println(numbers)

	colours:=[4]string{"Red","Blue", "Orange","Grapes"}
	fmt.Print(colours)

	ori:=[3]int{1,2,3}
	var copied *[3]int

	copied=&ori;
	



}