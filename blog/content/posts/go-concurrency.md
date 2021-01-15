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
[Refferences](#refferences)

### General

> If there's one thing most people know about Go, is that it is designed for concurrency [(1)](#1). \
> -- Andrew Gerrand -- 

#### Concurrent program

#### Process and thread

#### Concurrency vs parallel

#### Goroutine
The number of CPUs available simultaneously to executing goroutines is controlled by the GOMAXPROCS shell environment variable, whose default value is the number of CPU cores available ([2](#2))

#### Channel

#### Wait group




### Refferences
#### (1) 
https://golang.org/doc/faq#What_operations_are_atomic_What_about_mutexes
#### (2) 
https://golang.org/doc/faq#number_cpus
