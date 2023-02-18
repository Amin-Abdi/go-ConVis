package data

import "fmt"

func main6() {
	c1 := make(chan int)
	c2 := make(chan int)

	go gor1(c1)
	go gor2(c2)

	select {
	case msg1 := <-c1:
		fmt.Println("Received:", msg1)
	case msg2 := <-c2:
		fmt.Println("Received:", msg2)
	}

}

func gor1(a chan int) {
	a <- 99
}

func gor2(b chan int) {
	b <- 10
}
