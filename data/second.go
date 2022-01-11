package data

func main() {
	first := make(chan int)

	go ping(first)

	<-first

}

func ping(c chan int) {
	second := make(chan string)
	go pong(second)

	c <- 99
	second <- "Hello World"
	<-second

}

func pong(b chan string) {
	<-b
	b <- "variable 2"
}
