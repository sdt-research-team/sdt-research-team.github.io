package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// goroutineBasic()
	// waitGroup()
	sumConcurrency()
}

func goroutineBasic() {
	for i := 0; i < 10; i++ {
		timeNow := time.Now()
		go println("run", i, "start at", timeNow.Nanosecond(), "took", time.Since(timeNow), "(ns)")
	}

	time.Sleep(1 * time.Second) //Wait for all goroutine done before main terminated
}

func worker(id int, wg *sync.WaitGroup) {

	defer func() {
		fmt.Printf("waitgroup data: %+v", *wg)
		wg.Done()
		fmt.Println(runtime.NumGoroutine())
	}()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker %d done\n", id)

}

func waitGroup() {

	var wg = sync.WaitGroup{}

	for i := 1; i <= 5; i++ {
		// wg.Done()
		worker(i, &wg)
	}

	wg.Wait()
}

func addOne(n *int) {
	*n++
}

func sumConcurrency() {
	var result = 0

	for i := 0; i < 100; i++ {
		go addOne(&result)
	}

	for j := 0; j < 100; j++ {
		go addOne(&result)
	}

	fmt.Println("result", result)
}
