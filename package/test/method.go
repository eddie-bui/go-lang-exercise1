package test

import("fmt")

type sample struct{
	s string
	i int
}
func Hello(){
	fmt.Println("Hello World")
}
func (s sample) PrintName() {
	fmt.Printf("%v is %v",s.s,s.i)
}