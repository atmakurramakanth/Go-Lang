package main
  
import (
	"fmt"
	"runtime"
	"time"
)
  
func main() {
      
    switch day:=4; day{
       case 1:
       fmt.Println("Monday")
       case 2:
       fmt.Println("Tuesday")
       case 3:
       fmt.Println("Wednesday")
       case 4:
       fmt.Println("Thursday")
       case 5:
       fmt.Println("Friday")
       case 6:
       fmt.Println("Saturday")
       case 7:
       fmt.Println("Sunday")
       default: 
       fmt.Println("Invalid")
   }
   fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	t:=time.Now()
	switch{
	case t.Hour()<12:
		fmt.Println("goodMorning")
	case t.Hour()<17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("good evening")

	}
	// Deffered functions are pushed onto a stack When a function returns, its deferred calls are executed in last-in-first-out order.
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

     
}