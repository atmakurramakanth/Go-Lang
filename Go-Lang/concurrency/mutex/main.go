package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals=[]string{"test"}
var mut sync.Mutex   //pointer

var wg sync.WaitGroup  //pointer

func getStatusCode(endpoint string){
	defer wg.Done()  //this will pass a signal as done

	res,err:=http.Get(endpoint)

	if err!=nil{
		fmt.Println("Oops in endpoint")
	}else{
		//we keep a lock so that no other goroutine will try to write or append to signals array
		mut.Lock()
		signals=append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n",res.StatusCode,endpoint)
	}
}

func main(){

	websitelist :=[]string{
		"https://go.dev",
		"https://google.com",
		"https://github.com",
		"https://fb.com",
	}

	for _,web :=range websitelist{
		go getStatusCode(web)
		wg.Add(1)   //this will go on adding how many no of goroutines have completed
	}
	wg.Wait()   //this will make main func wait for signal from goroutine 
}
