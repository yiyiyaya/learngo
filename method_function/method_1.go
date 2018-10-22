package main

import "fmt"

func main() {
	var a N = 25
	a.value()
	p := &a
	p.value()
	(*p).value()
	a.pointer()
	p.pointer()
	

	//fmt.Println("v: %p,%v\n", &a, a)
}

type N int

func (n N) value() {
	n++
	fmt.Println("p: %p,%v\n", &n, n)
}

func (n *N) pointer() {
	(*n)++
	fmt.Println("h: %p,%v\n", n, *n)
}
