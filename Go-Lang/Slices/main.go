package main
import "fmt"
func main(){
	primes:=[6]int{2,3,5,7,11,13}
	var s []int =primes[1:4]
	fmt.Println(s)
	names := [4]string{
		"ram",
		"ramakanth",
		"reddy",
		"atmakur",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	//A slice literal is like an array literal without the length.
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	// Slice defaults
	z:= []int{2, 3, 5, 7, 11, 13}

	z = z[1:4]
	fmt.Println(z)

	z = z[:2]
	fmt.Println(z)

	z = z[1:]
	fmt.Println(z)
	//A slice has both a length and a capacity.
	//The length of a slice is the number of elements it contains.
	//The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	m := []int{2, 3, 5, 7, 11, 13}
	printSlice(m)

	// Slice the slice to give it zero length.
	m = m[:0]
	printSlice(s)
	m = m[:4]
	printSlice(m)
	m = m[2:]
	printSlice(m)
}

	

func printSlice(m []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(m),cap(m), m)

}

