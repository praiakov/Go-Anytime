package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // o canal foi fechado e drenado
			}
			squares <- x * x
		}

		close(squares)
	}()

	// Printer (na gorrotina principal)
	for x := range squares {
		fmt.Println(x)
	}
}
