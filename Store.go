package main

//This file is where the data structures for storing the types (i.e. channels, goroutines, functions) are stored

//type Goroutine struct {
//	name      string
//	origin    string
//	operation string
//	value     string
//}

type SendVale struct {
	origin      string
	value       string
	destination string
}
type ReceiveValue struct {
	origin      string
	destination string
}

//type Goroutine struct {
//	name string
//	origin string
//	params []string
//}
