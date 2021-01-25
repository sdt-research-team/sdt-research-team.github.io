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

#### Concurrent program
Today, modern systems are fast because they are using multiple cores. It is useful since we can split up bit process into smaller threads.

However, many basic tutorial only allow you to use one processor core because using multiple cores requires real threads. And it is totally hard because each thread will finish at a different time, and usually out of order. In reality, some tasks may require you to combine the result of each thread, and be impacted by the order in which you combine the results.

Instead of real threads, Go supports goroutines, which are lightweight threads. Go can handle goroutines by using tools from the language's stdlib itself. Moreover, Go can provide the right number of real threads to handle goroutines you spawn. 

A go routine is initally created with 2kb of stack size. Each function in go already has a check if more stack is needed or not and the stack can be copied to another region in memory with twice the original size. This makes goroutine very light on resources.

| Key                  | GoRoutine                                                              | Thread                                                                                       |
|----------------------|------------------------------------------------------------------------|----------------------------------------------------------------------------------------------|
| Managed By           | Goroutine methods are managed <br>by golang runtime.                   | Thread are managed by <br>operating systems.                                                 |
| Hardware dependency  | Goroutine are independent to hardware.                                 | Thread are dependent on hardware.                                                            |
| Communication Medium | Goroutines uses channels as communication <br>medium.                  | Thread have no easy communication medium.                                                    |
| Latency              | Goroutines can commuicate with other <br>routines with low latency.    | Thread as have no communication medium, <br>communicate with other thread with high latency. |
| ID                   | Goroutine does not have any thread local storage <br>and no unique id. | Thread have thread local storage and have <br>unique id.                                     |
| Scheduling           | Goroutines are co-operatively scheduled.                               | Threads are preemptively scheduled.                                                          |
| Startup              | Goroutines have faster startup time.                                   | Threads have slower startup time.                                                            |
| Stack                | Goroutines have growable segmented stacks.                             | Threads do not have growable segmented stacks.                                               |
#### Process and thread

#### Concurrency vs parallel

### Goroutine
The number of CPUs available simultaneously to executing goroutines is controlled by the GOMAXPROCS shell environment variable, whose default value is the number of CPU cores available ([2](#2))

### Wait group

### Channel

### Worker pool

### Rate Limit


### References
#### (1) 
https://golang.org/doc/faq#What_operations_are_atomic_What_about_mutexes
#### (2) 
https://golang.org/doc/faq#number_cpus
