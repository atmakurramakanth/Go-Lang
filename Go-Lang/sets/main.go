// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// place acts as a set of strings
// 	place := map[string]struct{}{}
	
// 	// We can add members to the set
// 	// by setting the value of each key to an
// 	// empty struct
// 	place["Hyd"] = struct{}{}
// 	place["Warangal"] = struct{}{}
// 	place["Mahabubnagar"] = struct{}{}
// 	place["Medchal"] = struct{}{}
	
// 	// Adding a new member to the set
// 	place["India"] = struct{}{}

// 	// Adding an existing to the set
// 	place["Mahabubnagar"] = struct{}{}
	
// 	// Removing a member from the set
// 	delete(place, "Warangal")
	
// 	fmt.Println(place)
// 	// iterating over the set using range operator
// 	for location := range place {
// 		fmt.Println(location)
// 	}


// 	// Adding methods to the set
// 	type locationSet map[string]struct{}

// 	// Adds an location to the set
// 	func (s LoclocationSet) add(location string){
// 		s[location]=struct(){}
// 	}

	
	
// }
package main
import "fmt"


type locationSet map[string]struct{}

// Adds an location to the set
func (s locationSet) add(location string) {
	s[location] = struct{}{}
}

// Removes an location from the set
func (s locationSet) remove(location string) {
	delete(s, location)
}

// Returns a boolean value describing if the location exists in the set
func (s locationSet) has(location string) bool {
	_, ok := s[location]
	return ok
}

func main() {
	// Initializing place as a new set
	place := locationSet{}
	
	// Adding members to the set
	place.add("Hyd")
	place.add("Warangal")
	place.add("Mahabubnagar")
	place.add("medchal")
	place.add("India")
	
	// Adding an existing member to the set again
	place.add("Mahabubnagar")
	
	// Removing members from the set
	place.remove("Warangal")
	
	fmt.Println(place)
	// map[Bear:{} Giraffe:{} Lion:{} Walrus:{}]
	
	// Checking the presence of elements in a set
	fmt.Println(place.has("Pakistan"))
	// false
	fmt.Println(place.has("India"))	
	// true
}