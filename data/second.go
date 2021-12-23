package data

func main() {
	first := make(chan int)
	third := make(chan int)

	go sender(first, "Hasicas", third)

	<-first
}

func sender(c chan int, k string, a chan int) {
	second := make(chan string)
	go pong(second)

	c <- 99
	second <- "Hello World"

}

func pong(b chan string) {
	<-b
}
