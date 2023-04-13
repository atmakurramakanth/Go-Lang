//concurrency is nothing but break the program into independently executing the tasks that should potentially run at the same time.
package main

import (
	"fmt"
	"time"
	
)
func main(){
	go count("sheep")
	count("fish")
}
func count(thing string) {
	for i:=1;true;i++ {
		fmt.Println(i,thing)
		time.Sleep(time.Millisecond*500)
	}
}