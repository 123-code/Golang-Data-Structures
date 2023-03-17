package main
import("fmt")

type Stack struct{
	elements[] int
}


func (s* Stack) Push(num int){
s.elements = append(s.elements,num);
}

func(s*Stack) Pop()[]int{
// remove last elements from the slice. 
	s.elements = s.elements[:len(s.elements)-1];
	return s.elements;


}

func main(){
	s:= new(Stack);
	s.Push(1);
	s.Push(2);
	s.Push(3);
	s.Push(4);
	fmt.Println(s);
	s.Pop();
	fmt.Println(s);
}
