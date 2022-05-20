package main   

import "fmt" 

type School interface{
	getName() string
	getAge()  int
}

type student struct{
	name string 
	age int
}

func (e *student) getName() string {
	return e.name
}

func (r *student) getAge() int {
	return r.age 
} 

func main(){
	var child student
	child.name = "munashe"
	child.age = 21

	fmt.Println(child.getAge())
	fmt.Println(child.getName())
}