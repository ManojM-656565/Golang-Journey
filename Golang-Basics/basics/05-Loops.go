package basics

import "fmt"

func Loops() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)

	}

	numbers:=[] int{1,2,3,4,5,6}
	for i,v :=range numbers {
		fmt.Printf("%v %v ",i,v)
	}


	for i:=range 10{
		fmt.Print(i);
	}
    i:=5
	for{
		//statenment block
		// server listeners , continous pooling , event loops 
        if i==7{
			break

		}
		i++
		// break or  return  to terminate the loop
	}

}