package main

import "fmt"

func main(){
	f1()

	f()
}

func f1(){
	fmt.Println("f1 - start")
	defer f2()
	fmt.Println("f1 - end")
}

func f2(){
	fmt.Println("f2-defferred")
}

func f(){
	for i:=0;i<5;i++{
		defer fmt.Printf("%d ", i)
	}
}