// make new
// new--- only allocates the memoryy 
// make - it allocates and init -non zeored storage

package main
import "fmt"
func main(){
	score :=make(map[string]int)
	score["ram"]=89
	fmt.Println(score)
	score["rahem"]=55
	score["ron"]=45
	score["sam"]=35
	fmt.Println(score)

	getRonScore:=score["ron"]
	fmt.Println(getRonScore)

	// delete function
	delete(score,"ron")
	fmt.Println(score)
	type Vertex struct {
		Lat, Long float64
	}
	
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)
}