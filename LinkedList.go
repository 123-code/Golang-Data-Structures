package main 
import ("fmt")

// LinkedList is a DS where values are stored by reference, using pointers

type node struct{
data int 
last *node
next *node 
}


type LinkedList struct{
head *node
length int

}

func(l*LinkedList) listnode(n*node){
second := l.head
l.head = n
l.head.next = second
l.length ++
}


func(l*LinkedList) DeleteNodeByValue(num int){
	if(num == l.head.data){ 
		l.head.data = l.head.next.data;
		l.length --;
	} else{
		// check if element exists in list
		for x :=0;x<l.length;x++{
if(l.head.next.data == num){
	l.head.data = l.head.next.data;
	l.length --;

}
fmt.Println("Node not found");

		}
	
	}

}


func(l LinkedList) PrintValues(){
	toPrint := l.head
	for l.length !=0{
		fmt.Printf("%d",toPrint.data)
		toPrint = toPrint.next
		l.length --;
	}
fmt.Println("\n");
}

func main(){
	mylist := new(LinkedList)
	node1 := &node{data: 1}
	node2 := &node{data:67}
	node3 := &node{data:89}
	node4:= &node{data:85}

mylist.listnode(node1);
mylist.listnode(node2);
mylist.listnode(node3);
mylist.listnode(node4);
mylist.PrintValues();
mylist.deleteNode(node2);
mylist.PrintValues();
//mylist.deleteNode();

}