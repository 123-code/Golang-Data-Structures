package main

import("fmt")

type Queue struct{
	elements[] int
}

func(s*Queue) queuein(num int){
	s.elements = append(s.elements,num);
}

func (s*Queue) queueout()[]int{
	//set s.elements to everything from index 1 to the end of the slice
s.elements = s.elements[1:len(s.elements)]
return s.elements;
}


func main(){
	queue:=new (Queue);

	queue.queuein(2);
	queue.queuein(3);
	queue.queuein(4);
	fmt.Println(queue);
	queue.queueout();
	fmt.Println(queue);
}