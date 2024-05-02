package main

import (
	"fmt"
	"time"
)

// func funcSatu(ch1 chan int) {
// 	ch1 <- 42
// }

// func funcDua(ch2 chan int) {
// 	ch2 <- 42
// }

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// go funcSatu(ch1)
	// go funcDua(ch2)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 42
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 100
	}()
	select {
	case nilai := <-ch1:
		fmt.Println(nilai)
	case nilai := <-ch2:
		fmt.Println(nilai)
	}

}
