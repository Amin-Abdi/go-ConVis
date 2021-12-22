package data

func main() {
	first := make(chan int)

	go sender(first)

	<-first

}

func sender(c chan int) {
	second := make(chan string)
	go pong(second)

	c <- 99
	<-second
}

func pong(b chan string) {
	b <- "Hello"
}
