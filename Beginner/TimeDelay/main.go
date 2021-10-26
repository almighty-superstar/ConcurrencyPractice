package main

import (
  "fmt"
  "time"
  "math/rand"
)

func main(){
	for i:=0;i<10;i++{
		go delayedPrint(i)
	}
	var input string
	fmt.Scanln(&input)
}

func delayedPrint(num int) {
	for i:=0;i<=10;i++{
		fmt.Println(num, ":",i)
		wait:=time.Duration(rand.Intn(200))*time.Millisecond
		time.Sleep(wait)
	}
}