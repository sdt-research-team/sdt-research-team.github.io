---
title: "Pointers"
date: 2020-12-19T09:57:01+07:00
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

## Escape Analysis
```go
// 
// 

// Sample program the mechanics of escape analysis.
package main

// user represents a user in the system.
type user struct {
	name  string
	email string
}

// main is the entry point for the application.
func main() {
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", u2)
}

// createUserV1 creates a user value and passed
// a copy back to the caller.
//go:noinline
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V1", &u)

	return u
}

// createUserV2 creates a user value and shares
// the value with the caller.
//go:noinline
func createUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V2", &u)

	return &u
}

/*
// See escape analysis and inlining decisions.

$ go build -gcflags -m=2
# github.com/ardanlabs/gotraining/topics/go/language/pointers/example4
./example4.go:24:6: cannot inline createUserV1: marked go:noinline
./example4.go:38:6: cannot inline createUserV2: marked go:noinline
./example4.go:14:6: cannot inline main: non-leaf function
./example4.go:30:16: createUserV1 &u does not escape
./example4.go:46:9: &u escapes to heap
./example4.go:46:9: 	from ~r0 (return) at ./example4.go:46:2
./example4.go:39:2: moved to heap: u
./example4.go:44:16: createUserV2 &u does not escape
./example4.go:18:16: main &u1 does not escape
./example4.go:18:27: main &u2 does not escape

// See the intermediate representation phase before
// generating the actual arch-specific assembly.

$ go build -gcflags -S
0x0021 00033 (/.../example4.go:15)	CALL	"".createUserV1(SB)
0x0026 00038 (/.../example4.go:15)	MOVQ	(SP), AX
0x002a 00042 (/.../example4.go:15)	MOVQ	8(SP), CX
0x002f 00047 (/.../example4.go:15)	MOVQ	16(SP), DX
0x0034 00052 (/.../example4.go:15)	MOVQ	24(SP), BX
0x0039 00057 (/.../example4.go:15)	MOVQ	AX, "".u1+40(SP)
0x003e 00062 (/.../example4.go:15)	MOVQ	CX, "".u1+48(SP)
0x0043 00067 (/.../example4.go:15)	MOVQ	DX, "".u1+56(SP)
0x0048 00072 (/.../example4.go:15)	MOVQ	BX, "".u1+64(SP)
0x004d 00077 (/.../example4.go:16)	PCDATA	$0, $1

// See bounds checking decisions.

go build -gcflags="-d=ssa/check_bce/debug=1"

// See the actual machine representation by using
// the disasembler.

$ go tool objdump -s main.main example4
TEXT main.main(SB) github.com/ardanlabs/gotraining/topics/go/language/pointers/example4/example4.go
example4.go:15	0x104aaf1		e8aa000000		CALL main.createUserV1(SB)
example4.go:15	0x104aaf6		488b0424		MOVQ 0(SP), AX
example4.go:15	0x104aafa		488b4c2408		MOVQ 0x8(SP), CX
example4.go:15	0x104aaff		488b542410		MOVQ 0x10(SP), DX
example4.go:15	0x104ab04		488b5c2418		MOVQ 0x18(SP), BX
example4.go:15	0x104ab09		4889442428		MOVQ AX, 0x28(SP)
example4.go:15	0x104ab0e		48894c2430		MOVQ CX, 0x30(SP)
example4.go:15	0x104ab13		4889542438		MOVQ DX, 0x38(SP)
example4.go:15	0x104ab18		48895c2440		MOVQ BX, 0x40(SP)

// See a list of the symbols in an artifact with
// annotations and size.

$ go tool nm example4
104aba0 T main.createUserV1
104ac70 T main.createUserV2
104ad50 T main.init
10be460 B main.initdone.
104aad0 T main.main
*/
```

* When a value could be referenced after the function that constructs the value returns.
* When the compiler determines a value is too large to fit on the stack.
* When the compiler doesn’t know the size of a value at compile time.
* When a value is decoupled through the use of function or interface values.

## Garbage Collection History

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

*- Toan Pham -*