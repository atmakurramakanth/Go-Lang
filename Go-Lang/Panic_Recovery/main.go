package main
import "fmt"


func ramakanth(){
	fmt.Println("ramakanth fun called")
	panicking()
	fmt.Println("ramakanth fun finished")
}
func panicking(){
	fmt.Println("It is called from ramakanth func")
	defer func() {
		if r:= recover(); r!=nil {
			fmt.Println("Recovery function called after panic happend")
		}
	}()
	panic("oh no")
	fmt.Println("panicking finshed")


}

func main(){
	fmt.Println("cascading panning")
	ramakanth()
	fmt.Println("end of the main function")
}