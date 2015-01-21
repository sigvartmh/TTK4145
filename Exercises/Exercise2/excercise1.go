// Go 1.2
// go run helloworld_go.go

package main

import (
	. "fmt" // Using '.' to avoid prefixing functions with their package names
	//   This is probably not a good idea for large projects...
	"runtime"
)

func Goroutine1(i_chan chan int, done chan bool) {
	for x := 0; x < 1000000; x++ {
		i := <-i_chan
		i++
		i_chan <- i
	}
	done <- true
}
func Goroutine2(i_chan chan int, done chan bool) {
	for x := 0; x < 1000000; x++ {
		i := <-i_chan
		i--
		i_chan <- i
	}
	done <- true
}

func main() {
	i_chan := make(chan int, 1)
	done := make(chan bool, 2)
	i_chan <- 0

	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	// Try doing the exercise both with and without it!
	go Goroutine1(i_chan, done) // This spawns someGoroutine() as a goroutine
	go Goroutine2(i_chan, done) // This spawns someGoroutine() as a goroutine
	// We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
	// We'll come back to using channels in Exercise 2. For now: Sleep.
	<-done
	<-done
	Printf("This is the value of i:%d\n", <-i_chan)
}
