package main

// import (
// 	"fmt"
// 	"time"
// )

// func functionGoroutineSatu(ch chan int) {
// 	ch <- 42
// 	fmt.Println("nilai dari goroutine pertama: ", ch)
// 	time.Sleep(time.Second)
// }

// func functionGoroutineDua(ch chan int, ch2 int) {
// 	ch <- ch2
// 	fmt.Println("adress dari goroutine kedua: ", ch)
// 	fmt.Println("nilai dari goroutine kedua: ", ch2)
// 	time.Sleep(time.Second)
// }

// func functionGoroutineTiga(ch3 int) {
// 	fmt.Println("nilai dari goroutine tiga: ", ch3)
// 	time.Sleep(time.Second)
// }

// func main() {
// 	ch := make(chan int)
// 	go functionGoroutineSatu(ch)
// 	ch2 := <-ch
// 	go functionGoroutineDua(ch, ch2)
// 	ch3 := <-ch
// 	go functionGoroutineTiga(ch3)
// 	time.Sleep(time.Second)
// }
