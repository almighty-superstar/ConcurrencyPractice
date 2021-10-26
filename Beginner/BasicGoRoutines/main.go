// package main

// import "fmt"

// func f(n int) {
//   for i := 0; i < 10; i++ {
//     fmt.Println(n, ":", i)
//   }
// }

// func main() {
//   go f(0)
//   var input string
//   fmt.Scanln(&input)
// }


package main

import "fmt"

func times3(){
	for i:=0; i<15;i++{
		fmt.Println(i*3)
	}
}

func main(){
	go times3()
	var slow string
	fmt.Scanln(&slow)
}