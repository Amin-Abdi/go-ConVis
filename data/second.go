package data

import "fmt"

func main() {
	first := make(chan int)
	third := make(chan string)

	go sender(first, "Helllo", third)
	go pong(third)

	<-first

}

func sender(c chan int, s string, rec chan string) {
	second := make(chan string)
	fmt.Println(second)
	c <- 99

	rec <- "Hey there!"
}

func pong(b chan string) {
	<-b
}
