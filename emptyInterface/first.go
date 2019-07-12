package main

import "fmt"

func main() {

	pack := Pack{}

	pack.Put("Hello", 123, true, "World", 4, Person{"Kazbek","Borash"}, 12.2, 123.1234, 'a')
	pack.PrintAll()
	pack.Drop("Hello")
	pack.Drop(true)
	pack.PrintAll()
}

type Person struct {
	name string
	lastName string
}

type Box interface {
	Put(a ...interface{}) int
	PrintAll()
	Drop(interface{})
}

type Pack struct {
	Types []interface{}
}

func (p *Pack) Put(a ...interface{}) {
	p.Types = append(p.Types, a...)
}

func (p *Pack) Drop(item interface{}) {
	for i, v := range p.Types {
		if v == item {
			fmt.Printf("Droped: %v\n", v)
			p.Types = append(p.Types[:i], p.Types[i+1:]...)
			return
		}
	}
	fmt.Printf("Not found: %v\n", item)
}

func (p *Pack) PrintAll() {
	for _,v := range p.Types {
		fmt.Printf("type of %v is %T\n", v,v)
	}
}


