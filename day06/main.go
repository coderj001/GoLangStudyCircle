// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello Go!!")
// 	main()
// }

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				fmt.Printf("goroutine %d: %d\n", n, j)
			}
		}(i)
	}

	wg.Wait()
}
