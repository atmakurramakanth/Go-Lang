package main

import (
	"fmt"
	
)
func add(x,y int)int{
	return x+y
}
func test(){
	fmt.Println("helo!")
}

func main(){
	ans:=add(5,6)
	fmt.Println(ans)
	x:=test // we can assign function to a variable
	x()

	// 
	t:=func(x int) int{
		return x*-1
	}(5)
	fmt.Println(t)
}