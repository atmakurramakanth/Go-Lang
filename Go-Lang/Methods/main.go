// Go does not have classes. However, you can define methods on types.
//A method is a function with a special receiver argument.
//Remember: a method is just a function with a receiver argument.
package main
import "fmt"
type Person struct{
	name string
	age int
}
func (p *Person) details(){
		p.name="new name"
		fmt.Println(p.name,p.age)
}
func main(){
	fmt.Println("Methods")
	p :=Person{
		name: "james",
		age: 25,
	}
	p.details()
	fmt.Printf("p = %+v\n",p)
}