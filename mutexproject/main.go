package main

import (
	"fmt"
	"os"
)

func fibonacci(n int) <-chan int {

	ch := make(chan int, n)

	for n!=0 {
		ch <- fib(n)
		n--
	}
	return ch
}

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func Write(n int) (chan int, chan int){
	ch := fibonacci(n)
	ch1 := make(chan int, n)
	ch2 := make(chan int, n)

	for i := 0; i < n; i++ {
		v := <-ch
		ch1 <- v
		ch2 <- v
	}

	return ch1, ch2
}

func main() {
	n := 4
	ch1 , ch2 := Write(n)

	f, err := os.Create("text.txt")

	if err != nil {
		fmt.Printf("Sorry, there is some problem writing to file...")
	}


	for i:=0; i<n; i++ {
		fmt.Printf("%d ", <-ch1)

		_, err := f.WriteString(fmt.Sprintf("%d ",<-ch2))
		if err != nil {
			fmt.Printf("Sorry can't write at index: %d", i)
		}
		f.Sync()
	}

	f.Close()

}
