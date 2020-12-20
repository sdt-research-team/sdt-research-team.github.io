---
title: "Pointers"
date: 2020-12-19T09:57:01+07:00
draft: false
tags: 
  - go
---

## Pointers

Pointers provide a way to share data across program boundaries. Having the ability to share and reference data with a pointer provides the benefit of efficiency. There is only one copy of the data and everyone can see it changing. The cost is that anyone can change the data which can cause side effects in running programs.

## Notes

* Use pointers to share data.
* Values in Go are always pass by value.
* "Value of", what's in the box. "Address of" ( **&** ), where is the box.
* The (*) operator declares a pointer variable and the "Value that the pointer points to".

## Escape Analysis

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