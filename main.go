package main

import ( "./Stacks")

func main(){
	s:= new(Stacks.Stack);
	s.Push(1);
	fmt.Println(s);
}