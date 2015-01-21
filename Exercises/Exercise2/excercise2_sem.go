// Go 1.2
// go run helloworld_go.go

package main

import (
	. "fmt" // Using '.' to avoid prefixing functions with their package names
	//   This is probably not a good idea for large projects...
	"runtime"
	//"time"
)

const MAX = 1000000

var i = 0

func Goroutine1(own chan bool, done chan bool) {
	for x := 0; x < MAX; x++ {
		<-own
		i++
		own <- true
	}
	done <- true
}
func Goroutine2(own chan bool, done chan bool) {
	for x := 0; x < MAX; x++ {
		<-own
		i--
		own <- true
	}
	done <- true
}

func main() {
	//i_chan := make(chan int, 1)
	//i_chan <- 0
	done := make(chan bool, 2)
	own := make(chan bool, 1)
	own <- true
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	// Try doing the exercise both with and without it!
	go Goroutine1(own, done) // This spawns someGoroutine() as a goroutine
	go Goroutine2(own, done) // This spawns someGoroutine() as a goroutine
	// We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
	// We'll come back to using channels in Exercise 2. For now: Sleep.
	<-done
	<-done

	Printf("This is the value of i:%d\n", i)
}
