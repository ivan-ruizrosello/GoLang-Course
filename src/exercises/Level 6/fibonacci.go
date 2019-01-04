package main

func main() {
	fib := fibonacci(7)

	println(fib)
}

func fibonacci (n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n - 1) + fibonacci(n -2)
}
