package data

import "fmt"

func main() {
	first := make(chan int)
	third := make(chan int)

	go sender(first, "Hasicas", third)
	go getStuff()

	<-first
}

func sender(c chan int, k string, a chan int) {
	c <- 99

	go getStuff()

}

func getStuff() {
	fmt.Println("AScas")
}
