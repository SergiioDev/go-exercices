package main

import "fmt"

func main() {
	c1 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c1 <- i
		}
		close(c1)
	}()

	for v := range c1 {
		fmt.Println(v)
	}

}
