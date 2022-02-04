package main

// to run use: go run problem3.go
import (
	"fmt"
	"time"
)

type FooVal struct {
	doFirst  chan bool
	doSecond chan bool
	doThird  chan bool
}

func (f *FooVal) first() {
	<-f.doFirst
	fmt.Print("first")
}

func (f *FooVal) second() {
	<-f.doSecond
	fmt.Print("second")
}

func (f *FooVal) third() {
	<-f.doThird
	fmt.Print("third")
}

func main() {
	executionOrder := []int{1, 3, 2}
	a := FooVal{make(chan bool), make(chan bool), make(chan bool)}
	go a.first()
	go a.second()
	go a.third()
	for _, next := range executionOrder {
		switch next {
		case 1:
			a.doFirst <- true
		case 2:
			a.doSecond <- true
		case 3:
			a.doThird <- true
		}
		time.Sleep(1 * time.Second)
	}
}
