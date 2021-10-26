package main 

import("fmt"
"time"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)
	go func() {
	  for i:=0;i<=12;i++ {
		chan1 <- "Cyborg likes Burgers and Beast Boy does not."
		time.Sleep(time.Millisecond * 250)
		if i==12{
			fmt.Println("Cyborg tried a burger and liked it.")
		}
	  }
	}()
  
	go func() {
	  for i:=0;i<=12;i++{
		chan2 <- "Beast Boy likes Burritos and Cyborg does not."
		time.Sleep(time.Millisecond * 280)
		if i==12{
			fmt.Println("Beast Boy tried a Burger and liked it.")
			fmt.Println("Cyborg and Beast Boy have both expanded their palates :).")
		}
	  }
	}()
  
	go func() {
	  for {
		select {
		case msg1 := <- chan1:
			fmt.Println(msg1)
		case msg2 := <- chan2:
			fmt.Println(msg2)
		}
	  }
	}()
		
	var input string
	fmt.Scanln(&input)
  }