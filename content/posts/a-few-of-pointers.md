---
title: "A Few Of Pointers"
date: 2020-12-18T10:57:01+07:00
author: "Pham Minh Toan"
draft: false
tags: 
  - go
---

## Pointers

> A *pointer* is a variable that points to the address of another variable. In computer science, pointers are a form of indirection, and indirection can be a powerful tool.

Pointers provide a way to share data across program boundaries. Having the ability to share and reference data with a pointer provides the benefit of efficiency. There is only one copy of the data and everyone can see it changing. The cost is that anyone can change the data which can cause side effects in running programs.

## Pass by Value
```go 
package main

func main() {

	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "value of" the count.
	increment(count)

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
}

// increment declares count as a pointer variable whose value is
// always an address and points to values of type int.
//go:noinline
func increment(inc int) {

	// Increment the "value of" inc.
	inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}
```

## Sharing data
```go 
package main

func main() {

	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")

	// Pass the "address of" count.
	increment(&count)

	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")
}

// increment declares count as a pointer variable whose value is
// always an address and points to values of type int.
//go:noinline
func increment(inc *int) {

	// Increment the "value of" count that the "pointer points to".
	*inc++

	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}
```

## Notes

* Use pointers to share data.
* Values in Go are always pass by value.
* "Value of", what's in the box. "Address of" ( **&** ), where is the box.
* The (*) operator declares a pointer variable and the "Value that the pointer points to".

## How do i know if my variable lives on the Stack or the Heap?
#### Does it matter ?
* Does not matter for the correctness of your program.
* Does affect the performance of your program.
* Anything on the heap is managed by the Gabage Collectior.
* The GC is very good, but it causes latency.
	* For the whole program, not just the part creating garbage.

#### Do you need to know ?
* If your program is not FAST ENOUGH...
* And you have benchmarks to prove it...
* And they show excessive heap allocations...
* Then maybe you should look into it.
* Optimize for correctness first, not performance.

### Why is the i/o reader interface use this way ? 
```go 
// io.Reader 
type Reader interface {
	Read(p []byte) (n int, err error)
}
```
Instead of 
```go 
type Reader interface {
	Read(n int) (b []byte, err error)
}
```

## Escape Analysis

1. The Stack
```go {linenos=table}
func main()  {
	n := 4
	n2 := square(n)
	println(n2)
}

func square(x int)int {
	return x*x
}
```

2. The Stack With Pointers 
```go {linenos=table} {linenos=table}
func main()  {
	n := 4
	inc(&n)
	println(n)
}

func inc(x *int){
	*x++
}
```

3. Returnning Pointers 
```go {linenos=table}
func main()  {
	n := answer()
	println(*n/2)
}

func answer() *int {
	x := 42
	return &x
}
```

> Sharing down typically stays on the stack.
> Sharing up typically escapes to the heap.
 [Link to the FAQ!](https://golang.org/doc/faq#stack_or_heap)

 #### Let's ask the compiler !
 ```go 
 go help build
 go tool compile -h
 go build -gcflags -m=1 // -m print optimization decisions
 ```

### When are values constructed on the heap ?
* When a value could be referenced after the function that constructs the value returns.
* When the compiler determines a value is too large to fit on the stack.
* When the compiler doesn’t know the size of a value at compile time.

### Which stays on the stack ?
```go 
func main() {
	b := read()
	// use b
}

func read() []byte {
	// return a new slice.
	b := make([]byte, 32)
	return b
}
```

```go 
func main() {
	b := make([]byte, 32)
	read(b)
	// use b
}

func read(b []byte) {
	// write into slice.
}
```
## Points to remember 
* Optimize for correctness, not performance.
* Go only puts function variables on the stack if it can prove a variable is not used after the function returns.
* Sharing down typically stays on the stack.
* Sharing up typically escapes to the heap.
* Ask the compiler to find out.
* Don't guess. Use the tooling

## Garbage Collection
* Tracking memory allocations in heap memory.
* Releasing allocations that are no longer needed.
* Keeping allocations that are still in-use.

## Pointers in disguise
Not all mutations require explicit use of a pointer. Go uses pointers behind the scenes for some of built-in collections.
* Maps are pointers
```go 
fun demolish(planets *map[string]string) <--  unnecessary pointer here
```
* Slices point at arrays
```go 
a := []int16{1,2,3}
b := a
fmt.Println(a,b) // [1,2,3] [1,2,3]

b[1] = 4
fmt.Println(a,b) // [1,4,3] [1,4,3]

fmt.Println(&a[0], &a[1]) // 0xc000090002 0xc000090004
fmt.Println(&b[0], &b[1]) // 0xc000090002 0xc000090004
```
## Enabling mutation
### Pointers as parameters
* Pointers are used to enable mutation across function and method boundaries.
Function and method parameters are passed by value in Go. That means functions always operate on a copy of passed agruments. When a pointer is passed to a function, the function receives a copy of the memory address. By dereferencing the memory address, a function can mutate the value a pointer point to.
```go 
type person struct {
  name, address string
  age           int
}

func birthday(p *person){
  p.age++
}

toan := person {
  name : "Toan Pham",
  address: "VN",
  age: 14
}

birthday(&toan)
fmt.Println("%+v\n", toan) 
// {name: "Toan Pham", address: "VN", age: 15}
```
Q: what age would toan be if the birthday(p person) function didn't use a pointer? 

### Pointer receivers
```go
type person struct {
  name string
  age int
}

func (p *person) birthday() {
  p.age++
}

an := &person{
  name: "An",
  age: 15,
}

an.birthday()
fmt.Printf("%+v\n", an)
// &{name: "An", age: 16}
```

#### You should use pointer receivers consistently. If some methods need pointer receivers, use pointer receivers for all methods of the type.

### Sumary
* Pointers store memory addresses.
* The address operator (&) provides the memory address of a variable.
* A pointer can be dereferenced (*) to access or modify the value it points to.
* Pointers are types declared with a preceding asterisk, such as *int.
* Use pointers to mutate values across function and method boundaries.
* Pointers are most useful with structures and arrays.
* Maps and slices use pointers behind the scenes.
* Interior pointer can point at fields inside structures without declaring those fields as pointers.
* Use pointers when it makes sense don't overuse them.

#### Pass Values

> Don't pass pointers as function arguments just to save a few bytes. If a function refers to its argument x only as *x throughout, then the argument shouldn't be a pointer. Common instances of this include passing a pointer to a string (*string) or a pointer to an interface value (*io.Reader). In both cases the value itself is a fixed size and can be passed directly. This advice does not apply to large structs, or even small structs that might grow.

*- Pham Minh Toan -*