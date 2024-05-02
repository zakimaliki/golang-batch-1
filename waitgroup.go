package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func myGoroutine(i int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Goroutine - Iteration", i)
// 	time.Sleep(time.Second)

// }

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 1; i <= 3; i++ {
// 		wg.Add(1)
// 		go myGoroutine(i, &wg)
// 	}
// 	wg.Wait()
// 	for i := 1; i <= 3; i++ {
// 		fmt.Println("main - Itration", i)
// 		time.Sleep(time.Second)
// 	}
// }
