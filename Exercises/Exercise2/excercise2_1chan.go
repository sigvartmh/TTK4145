// Go 1.2
// go run helloworld_go.go

package main

import (
	. "fmt" // Using '.' to avoid prefixing functions with their package names
	//   This is probably not a good idea for large projects...
	"runtime"
	"time"
)

const MAX = 1000000

func Goroutine1(i_chan chan int) {
	for x := 0; x < MAX-23; x++ {
		i := <-i_chan
		i++
		i_chan <- i
	}
}
func Goroutine2(i_chan chan int) {
	for x := 0; x < MAX; x++ {
		i := <-i_chan
		i--
		i_chan <- i
	}
}

func main() {
	i_chan := make(chan int, 1)
	i_chan <- 0
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	// Try doing the exercise both with and without it!
	go Goroutine1(i_chan) // This spawns someGoroutine() as a goroutine
	go Goroutine2(i_chan) // This spawns someGoroutine() as a goroutine
	time.Sleep(100 * time.Millisecond)
	Printf("This is the value of i:%d\n", <-i_chan)
}
