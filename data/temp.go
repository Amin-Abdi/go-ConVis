package data

import "fmt"

func main() {
	chan1 := make(chan int)

	go worker(chan1)

	<-chan1
	fmt.Println(<-chan1)
	fmt.Println(<-chan1)
	fmt.Println(<-chan1)

}

func worker(a chan int) {

	a <- 99
	a <- 24
	a <- 12
	a <- 88
}
