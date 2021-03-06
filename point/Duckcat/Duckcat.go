package Duckcat

import "fmt"

//import "fmt"



type Duck interface {
	Walk()
	Quack()
}

type Cat struct{}

func (c *Cat) Walk() {
	fmt.Println("catwalk")
}
func (c *Cat) Quack() {
	fmt.Println("meow")
}

func main() {
	//var c Duck = Cat{} 只有指针接收者
	var c Duck = &Cat{}
	c.Walk()
	c.Quack()
}