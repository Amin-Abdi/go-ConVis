package data

import "fmt"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go g1(c1)
	go g2(c2)

	select {
	case msg1 := <-c1:
		fmt.Println("Received:", msg1)
	case msg2 := <-c2:
		fmt.Println("Received:", msg2)
	}

}

func g1(a chan int) {
	a <- 99
}

func g2(b chan int) {
	b <- 10
}
