package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Go Routines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Go Routines\t", runtime.NumGoroutine())

	wg.Wait()
}
func x(s string) {
	for x := 0; x < int(math.Pow(10, 6)); x++ {
		fmt.Println(s, x)
	}
}

func foo() {
	x("foo:")
	wg.Done()
}
func bar() {
	x("bar:")
}
