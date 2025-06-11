package basics;

import "fmt"

const GRAVITY=9.81
const PI=3.14
func Constants(){
	const(
		monday=1
		tuesday=3 //untyped const
		wednesday int=3 //typed const

	)
	fmt.Print(monday)
}