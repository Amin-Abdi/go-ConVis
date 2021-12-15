package data

import "fmt"

func main() {
	first := make(chan int)
	third := make(chan string)

	go sender(first, "Helllo", third)

	<-first
	
}

func sender(c chan int, s string, b chan string) {
	second := make(chan string)
	fmt.Println(second)
	c <- 99
}
