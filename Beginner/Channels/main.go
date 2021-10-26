package main

import (
  "fmt"
  "time"
  "math/rand"

)

func pinger(c chan string) {
	sites:=[]string{"facebook.com","google.com","apple.com","amazon.com","netflix.com","microsoft.com"}
	var choice string
	for i := 0;i<=15; i++ {
		choice=sites[rand.Intn(len(sites))]
		c <- choice
  }
}

func printer(c chan string) {
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func main() {
  var c chan string = make(chan string)

  go pinger(c)
  go printer(c)

  var input string
  fmt.Scanln(&input)
}