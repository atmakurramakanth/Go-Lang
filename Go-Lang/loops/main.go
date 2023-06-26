package main

import "fmt"
func main(){
	sum:=0
	for i:=0;i<10;i++{
		sum+=i
	}
	fmt.Println(sum)
	add:= 1
	for ; add < 1000; {
		add+= add
	}
	fmt.Println(add)
	s:= 1
	for s < 1000 {
		s += s
	}
	fmt.Println(s)
	// infinte loop
	// for {
	// }
	

}
