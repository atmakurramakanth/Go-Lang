package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	// struct fields are accessed using dot
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	//Struct fields can be accessed through a struct pointer.
	a := Vertex{1, 2}
	p := &a
	p.X = 1e9
	fmt.Println(a)

}