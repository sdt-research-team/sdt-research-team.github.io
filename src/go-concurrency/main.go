package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// goroutineBasic()
	// waitGroup()
	// sumConcurrency()
	waitGroupWithoutGoroutine()
}

func goroutineBasic() {
	for i := 0; i < 10; i++ {
		timeNow := time.Now()
		go println("run", i, "start at", timeNow.Nanosecond(), "took", time.Since(timeNow), "(ns)")
	}

	time.Sleep(1 * time.Second) //Wait for all goroutine done before main terminated
}

func _worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func waitGroup() {

	var wg = sync.WaitGroup{}

	for i := 1; i <= 5; i++ {
		// wg.Done()
		_worker(i, &wg)
	}

	wg.Wait()
}

func waitGroupWithoutGoroutine() {
	var wg = sync.WaitGroup{}
	wg.Add(5)

	for i := 1; i <= 5; i++ {
		fmt.Println("Doing task #", i)
		wg.Done()
	}

	wg.Wait()
}

func _addOne(n *int) {
	*n++
}

func sumConcurrency() {
	var result = 0

	for i := 0; i < 100; i++ {
		go _addOne(&result)
	}

	for j := 0; j < 100; j++ {
		go _addOne(&result)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("result", result)
}
