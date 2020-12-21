---
title: "Control Structures"
date: 2020-12-19T09:57:01+07:00
author: "Huynh Hong An"
draft: false
tags: 
  - go
---

# Control Structures

### Table of content

[General](#general)\
[If](#if)\
[Redeclaration And Reassignment](#redeclaration-and-reassignment)\
[For](#for)\
[Switch](#switch)

### General

### If
Basic usage
```go
if x > 0 {
    return y
}
```
_Note: Open & close parentheses "()" are not required_

With initialization
```go
if err := Foo(); err != nil {
    return err
}
```

### Redeclaration and reassignment
- Appears in **Short variable declarations** with at least one other variable that is created by the declaration.
Example:
```go 
func Foo() (int, int) {
	return 1,2
}

func main() {
	var a = 10
	if true {
		a, b := Foo()// This is equal to "a, b := 1, 2"
		fmt.Println("Second a: ", a)
		_ = b
	}

	fmt.Println("First a:", a)
}
```
The seccond `a` is in defferent scope from the first `a`
Output
```bash
Second a: 1
First a: 10
```

### For

### Switch 
#### Expression switch
- As `if`, `switch` also support initial statment before condition.
- If `switch` has no expression, it switches on `true`.

Example
```go
switch {
    case sound == "moo": //if (sound == "moo") == true { return "cow" }
        return "cow"
    case sound == "meow":
        return "cat"
    default: 
        return "human"
}
```
Similar to this syntax
```go
switch sound {
    case "moo":
        return "cow"
    case "meow":
        return "cat" 
    default:
        return "human"
}
```
- If `case` is an untyped constant, then it will convert to match data type in `switch` statement.

Example
```go
var myInt16 int8 = 127
	switch myInt16 {
	case 127:
		fmt.Println("First case")
	case 128: //Compile error: constant 128 overflows int8
		fmt.Println("Second case")
	}
```
More example, how about now? [(1)](#(1))
```go
var myInterface interface{} = int8(127)

switch myInterface {
case 127:
    fmt.Println("First case")
case 1000: // Is overflows int8?
    fmt.Println("Second case")
}
```

- By default, `break` is automatically added to the end of `case`. To make a multiple case, list them in same `case` and seperated by `,` (EX: `case "moo", "low":`).

- `fallthrough` statment allows next `case` execute next statement without taking care of next condition.
```go
//Example of fallthrough
var flavor = "strawberry"
switch flavor {
    case "strawberry":
        fmt.Println(flavor, "is my favorite!")
        fallthrough
    case "vanilla", "chocolate":
        fmt.Println(flavor, "is great!")
    default:
        fmt.Println("I've never tried", flav, "before")
}
// strawberry is my favorite!
// strawberry is great!
```

- Switch cases execution order: top to bottom, left to right.

```go
// Foo prints and returns n.
func Foo(n int) int {
    fmt.Println(n)
    return n
}

func main() {
    switch Foo(2) {
    default:
		fmt.Println("Default")
    case Foo(1), Foo(2), Foo(3):
        fmt.Println("First case")
        fallthrough
    case Foo(4):
        fmt.Println("Second case")
    }
}
```

❓ How to break a loop while in a `case`? [(2)](#(2))

#### Type switch

- Basic structure
```go
switch i := x.(type) { // x must be an interface
case nil:
	printString("x is nil")                // type of i is type of x (interface{})
case int:
	printInt(i)                            // type of i is int
case float64:
	printFloat64(i)                        // type of i is float64
case func(int) float64:
	printFunction(i)                       // type of i is func(int) float64
case bool, string:
	printString("type is bool or string")  // type of i is type of x (interface{})
default:
	printString("don't know the type")     // type of i is type of x (interface{})
}
```
- `fallthough` is not allowed to use in **type switch**

❓ `switch` vs `if & else` [(3)](#(3))

### Summary


### Refferences
#### (1)
```go
var myInterface interface{} = int8(127)

switch myInterface {
case 127:
    fmt.Println("First case")
case 1000: // Is overflows int8?
    fmt.Println("Second case")
}
```
Output
```bash
```
*Yes, seriously, nothing!*

#### (2)
Using label for loop
```go
	var animal string
	var sound = "moo"

Loop:
	for i := 0; i < 10; i ++ {
		fmt.Println("gau gau")
		switch sound {
		case "moo":
			animal = "cow"
			break Loop
		case "meow":
			animal = "cat"
		}
	}
```

#### (3)
https://www.geeksforgeeks.org/switch-vs-else/