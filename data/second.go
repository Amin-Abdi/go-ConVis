package data

func main() {
	first := make(chan int)

	go go1(first)

	<-first

}

func go1(c chan int) {
	second := make(chan string)
	go go2(second)

	c <- 99
	<-second

}

func go2(b chan string) {
	//<-b
	b <- "variable 2"
}
