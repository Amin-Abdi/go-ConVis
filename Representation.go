package main

//This file is where the data structures for storing the types (i.e. channels, goroutines, functions) are stored

type Creation struct {
	TypeOp string
	Name   string
	Parent string
}

type SendRec struct {
	TypeOp      string
	Origin      string
	Destination string
	Value       string
}
