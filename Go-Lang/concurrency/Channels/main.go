package main

import (
	"fmt"
	"sync"
)

//error: all goroutines are asleep - deadlock!

// func foo(c chan int,someValue int){
// 	c <- someValue*5
// }

// func main(){
// 	fooVal:=make(chan int)

// 	// go foo(fooVal,5)
// 	// go foo(fooVal,3)

// 	// v1,v2:= <-fooVal,<-fooVal

// 	// let's ,make iterative
// 	for i:=0;i<10;i++ {
// 		go foo(fooVal,i)
// 	}
// 	for item:=range fooVal{
// 		fmt.Println(item)
// 	}

// }

// to resolve the error

var wg sync.WaitGroup

func foo(c chan int,someValue int){
	defer wg.Done()
	c <- someValue*5
}

func main(){
	fooVal:=make(chan int,10) // we need to add buffer to channel

	for i:=0;i<10;i++ {
		wg.Add(1)
		go foo(fooVal,i)
	}
	wg.Wait()
	close(fooVal)
	for item:=range fooVal{
		fmt.Println(item)
	}
	
}

