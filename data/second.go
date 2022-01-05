package data

func main() {
	first := make(chan int)

	go ping(first, "Hasicas")

	<-first

}

func ping(c chan int, k string) {
	second := make(chan string)
	go pong(second)

	c <- 99
	second <- "Hello World"

}

func pong(b chan string) {
	<-b
}
