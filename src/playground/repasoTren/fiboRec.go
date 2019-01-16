package repasoTren

import "fmt"

func main() {
	for i := 0; i < 15; i ++ {
		fmt.Println(fibo(i))
	}
}

func fibo (n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fibo(n-1) + fibo(n-2)
}
