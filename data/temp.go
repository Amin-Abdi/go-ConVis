package data

func main() {
	con := make(chan int)

	go temp1(con)
}

func temp1(a chan int) {
	a <- 99
}
