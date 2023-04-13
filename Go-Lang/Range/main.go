package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
// (in Range) for index,value:= pow{
//}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	for _,x:= range pow{ // if we keep index value without specifying it gives error
		fmt.Printf("the value is %d",x)
	}
}