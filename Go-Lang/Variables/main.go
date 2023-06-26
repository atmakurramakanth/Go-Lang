package main

import "fmt"

var c, python, java bool // variable declaration
var a, j int = 1, 2        // variables with initalizers

func main() {
	var i int
	k := 3
	fmt.Println(i, c, python, java)
	var c, python, java = true, false, "no!"
	fmt.Println(k,a, j, c, python, java)
}