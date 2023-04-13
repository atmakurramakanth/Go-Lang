package main

import (
	"fmt"
	"math"
	"math/rand"
)
func add(x int, y int) int{   // x int, y int ==>  x,y int
	return x+y
}
func swap(x,y string)(string,string){
	return y,x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(add(42,13))
	a,b:=swap("hello","world")
	fmt.Println(a,b)
	fmt.Println(split(17))
}