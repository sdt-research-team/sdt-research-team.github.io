---
title: "Go Concurrency"
date: 2021-01-14T09:57:01+07:00
author: "Pham Minh Toan, Huynh Hong An, Pham Quoc Dat"
draft: false
tags:
  - go
  - operating system
---

### Table of content

[General](#general)\
[Goroutine](#goroutine)\
[Wait group](#wait-group)\
[Channel](#channel)\
[Worker pool](#worker-pool)\
[Rate limit](#rate-limit)\
[References](#references)

### General

> If there's one thing most people know about Go, is that it is designed for concurrency [(1)](#1). \
> -- Andrew Gerrand --
**concurrency and parallelism**
Concurrency is the ability of a computer to deal with multiple things at once. And, concept of running multiple things in parallel known as parallelism

#### Concurrent program

Today, modern systems are fast because they are using multiple cores. It is useful since we can split up bit process into smaller threads.

However, many basic tutorial only allow you to use one processor core because using multiple cores requires real threads. And it is totally hard because each thread will finish at a different time, and usually out of order. In reality, some tasks may require you to combine the result of each thread, and be impacted by the order in which you combine the results.

Instead of real threads, Go supports goroutines, which are lightweight threads. Go can handle goroutines by using tools from the language's stdlib itself. Moreover, Go can provide the right number of real threads to handle goroutines you spawn.

At any point in time, one thread will be executing one goroutine and if that goroutine is blocked, then it will be swapped out for another goroutine that will execute on that thread. It looks like **thread scheduling** but handled by **Go runtime** and this is much faster.

A go routine is initally created with 2kb of stack size. Each function in go already has a check if more stack is needed or not and the stack can be copied to another region in memory with twice the original size. This makes goroutine very light on resources.

| Key                  | GoRoutine                                                              | Thread                                                                                       |
| -------------------- | ---------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| Managed By           | Goroutine methods are managed <br>by golang runtime.                   | Thread are managed by <br>operating systems.                                                 |
| Hardware dependency  | Goroutine are independent to hardware.                                 | Thread are dependent on hardware.                                                            |
| Communication Medium | Goroutines uses channels as communication <br>medium.                  | Thread have no easy communication medium.                                                    |
| Latency              | Goroutines can commuicate with other <br>routines with low latency.    | Thread as have no communication medium, <br>communicate with other thread with high latency. |
| ID                   | Goroutine does not have any thread local storage <br>and no unique id. | Thread have thread local storage and have <br>unique id.                                     |
| Scheduling           | Goroutines are co-operatively scheduled.                               | Threads are preemptively scheduled.                                                          |
| Startup              | Goroutines have faster startup time.                                   | Threads have slower startup time.                                                            |
| Stack                | Goroutines have growable segmented stacks.                             | Threads do not have growable segmented stacks.                                               |


Imagine, you have a web server, this server is handling 1000 requests per second. If an OS thread consume 1MB stack size per thread, that means it takes 1GB of RAM for that traffic
#### Process and thread

#### Concurrency vs parallel

### Goroutine

The number of CPUs available simultaneously to executing goroutines is controlled by the GOMAXPROCS shell environment variable, whose default value is the number of CPU cores available ([2](#2))

#### Basic syntax
Goroutine in a loop
```go
func main() {
	for i := 0; i < 10; i++ {
		go println(i)
	}
  //Wait for all goroutine done before main terminated
	time.Sleep(1 * time.Second) 
}

```
Output
```go
1
9
2
3
4
5
6
8
0
7
```

#### Goroutine under the hood
- When initialize a goroutine, it uses stacks. It provides a few KBs of stack RAM, which is almost always enough. This amount of RAM is based on strategy of Golang, for example ([3](#3)):
  - Go 1.2: goroutine stack has been increased from 4Kb to 8Kb.
  - Go 1.4: goroutine stack has decreased from 8Kb to 2Kb.

- When **goroutine size overflow** the init size. Runtime  Golang has some methods for stack management:
  - Segmented stacks
    - Example, we have stack S1 is currently init by a goroutine G1. If the stack S1 is almost full, a call will force a new stack chunk to be allocated.\
    ![stack-segment-01.png](https://raw.githubusercontent.com/sdt-research-team/sdt-research-team.github.io/main/blog/static/images/stack-segment-01.png)
    - **Problem**: This approach can become hell if the stack is almost full in a loop, it will create multiple useless segments.
  - Stack copying
    - Example, we have stack S1 is currently init by a goroutine G1. If the stack S1 is almost full, it will allocate another stack S2 and copy all S1 data to S2 (S1 will be cleaning up later by garbage collector) \
    ![stack-copy-01.png](https://raw.githubusercontent.com/sdt-research-team/sdt-research-team.github.io/main/blog/static/images/stack-copy-01.png)

#### Why goroutines instead of threads?


### Wait group

#### Basic syntax
```go
func worker(id int, wg *sync.WaitGroup) {

    defer wg.Done() //Decrease counter by one

    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1) //Increase counter by one
        go worker(i, &wg)
    }

    wg.Wait() //Wait for counter until zero
}
```

#### WaitGroup and concurrency
WaitGroup is independent with concurrency programing. Because you can use it without any goroutines. Example:
```go
func waitGroupWithoutGoroutine() {
	var wg = sync.WaitGroup{}
	wg.Add(5)

	for i := 1; i <= 5; i++ {
		fmt.Println("Doing task #", i)
		wg.Done()
	}

	wg.Wait()
}
```

But WaitGroup is developed for concurrency. The type **WaitGroup** is in package *sync* which is provides basic synchronization primitives (thread or process synchronization frameworks). 

#### WaitGroup specifications: 
- Should not be copied. When passing a wait group variable to a goroutine function, it should be passed by pointer reference.
- Use stack pointer to store counter value.




### Channel

### Worker pool

### Rate Limit

### References

#### (1)
https://golang.org/doc/faq#What_operations_are_atomic_What_about_mutexes

#### (2)
https://golang.org/doc/faq#number_cpus

#### (3)
https://medium.com/a-journey-with-go/go-how-does-the-goroutine-stack-size-evolve-447fc02085e5