# [Reference](https://golangbot.com/learn-golang-series/)

[I - Variables, Types and Constants](#1)
[Variables](#Variables)
[Types](#Types)
[Constants](#Constants)
[II - Functions and Packages](#2)
[Functions](#Functions)
[Go Packages](#Go-Packages)
[III - Conditional Statements and Loops](#3)
[If else statement](#if-statement)
[Loops](#Loops)
[Switch Statement](#switch)
[Arrays and Slices](#arrays-and-slices)
[Variadic Functions](#variadic-functions)

# I - Variables, Types and Constants {#1}

## [Variables](https://golangbot.com/variables/) {#Variables}

- Declaring a single variable: `var name type`
- Declaring a variable with an initial value: `var name type = initialvalue`
- Type inference (Tự suy luận type dựa trên initial value) `var name = initialvalue`
- Multiple variable declaration: `var name1, name2 type = initialvalue1, initialvalue2`
  There might be cases where we would want to declare variables belonging to different types in a single statement. The syntax for doing that is

```
var (
      name1 = initialvalue1
      name2 = initialvalue2
)
```

Example:

```go
package main

import "fmt"

func main() {
    var (
        name   = "naveen"
        age    = 29
        height int
    )
    fmt.Println("my name is", name, ", age is", age, "and height is", height)
}
```

- Short hand declaration (khỏi dùng var mà dùng := operator): `name := initialvalue`

## [Types](https://golangbot.com/types/) {#Types}

the basic types available in Go:

1. bool
2. Numeric Types

- int8, int16, int32, int64, int
- uint8, uint16, uint32, uint64, uint
- float32, float64
- complex64, complex128
- byte
- rune

3. string

#### Signed integers

**int8:** represents 8 bit signed integers
**size:** 8 bits
**range:** -128 to 127

**int16:** represents 16 bit signed integers
**size:** 16 bits
**range:** -32768 to 32767

**int32:** represents 32 bit signed integers
**size:** 32 bits
**range:** -2147483648 to 2147483647

**int64:** represents 64 bit signed integers
**size:** 64 bits
**range:** -9223372036854775808 to 9223372036854775807

**int:** represents 32 or 64 bit integers depending on the underlying platform. You should generally be using int to represent integers unless there is a need to use a specific sized integer.
**size:** 32 bits in 32 bit systems and 64 bit in 64 bit systems.
**range:** -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems

#### Unsigned integers

**uint8:** represents 8 bit unsigned integers
**size:** 8 bits
**range:** 0 to 255

**uint16:** represents 16 bit unsigned integers
**size:** 16 bits
**range:** 0 to 65535

**uint32:** represents 32 bit unsigned integers
**size:** 32 bits
**range:** 0 to 4294967295

**uint64:** represents 64 bit unsigned integers
**size:** 64 bits
**range:** 0 to 18446744073709551615

**uint:** represents 32 or 64 bit unsigned integers depending on the underlying platform.
**size:** 32 bits in 32 bit systems and 64 bits in 64 bit systems.
**range:** 0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems

#### Floating point types

**float32:** 32 bit floating point numbers
**float64:** 64 bit floating point numbers

#### Complex types

**complex64**: complex numbers which have float32 real and imaginary parts
**complex128**: complex numbers with float64 real and imaginary parts

- The builtin function complex is used to construct a complex number with real and imaginary parts. The complex function has the following definition
  `func complex(r, i FloatType) ComplexType`
- It takes a real and imaginary part as a parameter and returns a complex type. Both the real and imaginary parts must be of the same type. ie either float32 or float64. If both the real and imaginary parts are float32, this function returns a complex value of type complex64. If both the real and imaginary parts are of type float64, this function returns a complex value of type complex128

- Complex numbers can also be created using the shorthand syntax
  `c := 6 + 7i`

#### Other numeric types

byte is an alias of uint8
rune is an alias of int32

#### Type Conversion

Go is very strict about explicit typing. There is no automatic type promotion or conversion

```go
package main

import (
    "fmt"
)

func main() {
    i := 55      //int
    j := 67.8    //float64
    sum := i + j //int + float64 not allowed
    fmt.Println(sum)
}
```

to fix the above error. `T(v) is the syntax to convert a value v to type T`

```go
package main

import (
    "fmt"
)

func main() {
    i := 55      //int
    j := 67.8    //float64
    sum := i + int(j) //j is converted to int
    fmt.Println(sum)
}
```

## [Constants](https://golangbot.com/constants/) {#Constants}

- Declaring a constant: The keyword `const` is used to declare a constant
- Declaring a group of constants:

```go
package main

import (
    "fmt"
)

func main() {
    const (
        name = "John"
        age = 50
        country = "Canada"
    )
    fmt.Println(name)
    fmt.Println(age)
    fmt.Println(country)

}
```

# II - Functions and Packages {#2}

## [Functions](https://golangbot.com/functions/) {#Functions}

- Function declaration:

```go
func functionname(parametername type) returntype {
 //function body
}
```

The parameters and return type are optional in a function

```go
func functionname() {
}
```

If consecutive parameters are of the same type, we can avoid writing the type each time and it is enough to be written once at the end

```go
func calculateBill(price, no int) int {
    var totalPrice = price * no
    return totalPrice
}
```

- Multiple return values:
  It is possible to return multiple values from a function

```go
package main

import (
    "fmt"
)

func rectProps(length, width float64)(float64, float64) {
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}

func main() {
     area, perimeter := rectProps(10.8, 5.6)
    fmt.Printf("Area %f Perimeter %f", area, perimeter)
}
```

- Named return values:
  It is possible to return named values from a function. If a return value is named, it can be considered as being declared as a variable in the first line of the function.

```go
func rectProps(length, width float64)(area, perimeter float64) {
    area = length * width
    perimeter = (length + width) * 2
    return //no explicit return value
}
```

- Blank Identifier:
  \_ is known as the blank identifier in Go. It can be used in place of any value of any type.

```go
package main

import (
    "fmt"
)

func rectProps(length, width float64) (float64, float64) {
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}
func main() {
    area, _ := rectProps(10.8, 5.6) // perimeter is discarded
    fmt.Printf("Area %f ", area)
}
```

In line no. 13 we use only the area and the \_ identifier is used to discard the perimeter.

## [Go Packages](https://golangbot.com/go-packages/) {#Go-Packages}

- **Packages are used to organize Go source code for better reusability and readability. Packages are a collection of Go sources files that reside in the same directory. Packages provide code compartmentalization and hence it becomes easy to maintain Go projects.**

main function and main package:

- Every executable Go application must contain the main function. This function is the entry point for execution. The main function should reside in the main package.

`package packagename` **specifies that a particular source file belongs to package** `packagename`. **This should be the first line of every go source file.**

#### Go Module

**A Go Module is nothing but a collection of Go packages**. Now this question might come to your mind. Why do we need Go modules to create a custom package? **The answer is the import path for the custom package we create is derived from the name of the go module**

#### Creating a Go module

`go mod init learnpackage`
The above command will create a file named `go.mod`. The following will be the contents of the file.

```
module learnpackage

go 1.13
```

The line `module learnpackage` specifies that the module's name is `learnpackage`. As we mentioned earlier, `learnpackage` will be the base path to import any package created inside this module. The line `go 1.13` specifies that the files in this module use go version `1.13`.

#### Exported Names

- Any variable or function which starts with a **capital letter** are exported names in go. **Only exported functions and variables can be accessed from other packages.**
  `=> if you want to access a function outside of a package, it should be capitalized.`

#### init function

Each package in Go can contain an init function. The init function must not have any return type and it must not have any parameters. The init function cannot be called explicitly in our source code. It will be called automatically when the package is initialized. The init function has the following syntax

```go
func init() {
}
```

**The order of initialization of a package is as follows**

```
1. Package level variables are initialised first
2. init function is called next. A package can have multiple init functions (either in a single file or distributed across multiple files) and they are called in the order in which they are presented to the compiler.

If a package imports other packages, the imported packages are initialized first.
A package will be initialized only once even if it is imported from multiple packages.
```

#### Use of blank identifier

It is illegal in Go to import a package and not to use it anywhere in the code. The compiler will complain if you do so. The reason for this is to avoid bloating of unused packages which will significantly increase the compilation time.

```go
package main

import (
        "learnpackage/simpleinterest"
)

func main() {

}
```

The above program will error

But it is quite common to import packages when the application is under active development and use them somewhere in the code later if not now. The \_ blank identifier saves us in those situations.
The error in the above program can be silenced by the following code,

```go
package main

import (
        "learnpackage/simpleinterest"
)

var _ = simpleinterest.Calculate

func main() {

}
```

The line var \_ = simpleinterest.Calculate mutes the error.

Sometimes we need to import a package just to make sure the initialization takes place even though we do not need to use any function or variable from the package. For example, we might need to ensure that the `init` function of package is called even though we plan not to use that package anywhere
The \_ blank identifier can be used in this case too as shown below.

```go
package main

import (
    _ "learnpackage/simpleinterest"
)

func main() {

}
```

# III - Conditional Statements and Loops {#3}

## [If else statement](https://golangbot.com/if-statement/) {#if-statement}

- If statement syntax:

```go
if condition {
}
```

Unlike in other languages like C, **the braces { } are mandatory** even if there is only one line of code between the braces{ }.

- If else statement

```go
if condition {
} else {
}
```

The if statement has an optional else construct which will be executed if the condition in the if statement evaluates to false.

- If ... else if ... else statement

```go
if condition1 {
...
} else if condition2 {
...
} else {
...
}
```

- If with assignment

There is one more variant of if which includes an optional shorthand assignment statement that is executed before the condition is evaluated. Its syntax is

```go
if assignment-statement; condition {
}
```

In the above snippet, `assignment-statement` is first executed before the condition is evaluated.

Let's rewrite the program which finds whether the number is even or odd using the above syntax.

```go
package main

import (
    "fmt"
)

func main() {
    if num := 10; num % 2 == 0 { //checks if number is even
        fmt.Println(num,"is even")
    }  else {
        fmt.Println(num,"is odd")
    }
}
```

In the above program num is initialized in the if statement in line no. 8. **_One thing to be noted is that num is available only for access from inside the if and else. i.e. the scope of num is limited to the if else blocks. If we try to access num from outside the if or else, the compiler will complain_**

**The else statement should start in the same line after the closing curly brace } of the if statement. If not the compiler will complain.**
The reason is because of the way Go inserts semicolons automatically. You can read about the semicolon insertion rule here https://golang.org/ref/spec#Semicolons.

## [Loops](https://golangbot.com/loops/) {#Loops}

for is the only loop available in Go. Go doesn't have while or do while loops which are present in other languages like C.

- for loop syntax:

```go
for initialisation; condition; post {
}
```

The initialisation statement will be executed only once. After the loop is initialised, the condition will be checked. If the condition evaluates to true, the body of the loop inside the { } will be executed followed by the post statement. The post statement will be executed after each successful iteration of the loop. After the post statement is executed, the condition will be rechecked. If it's true, the loop will continue executing, else the for loop terminates.
All the three components namely initialisation, condition and post are optional in Go. Let's look at an example to understand for loop better.

- break
  The `break` statement is used to terminate the for loop abruptly before it finishes its normal execution and move the control to the line of code just after the for loop.
- continue
  The `continue` statement is used to skip the current iteration of the for loop. All code present in a for loop after the continue statement will not be executed for the current iteration. The loop will move on to the next iteration.
- Nested for loops
  A `for` loop which has another `for` loop inside it is called a nested for loop.
- Labels
  Labels can be used to break the outer loop from inside the inner for loop.

```go
package main

import (
    "fmt"
)

func main() {
outer:
    for i := 0; i < 3; i++ {
        for j := 1; j < 4; j++ {
            fmt.Printf("i = %d , j = %d\n", i, j)
            if i == j {
                break outer
            }
        }

    }
}
```

- infinite loop
  The syntax for creating an infinite loop is,

```go
for {
}
```

```go
package main

import "fmt"

func main() {
    for {
        fmt.Println("Hello World")
    }
}
```

## [Switch Statement](https://golangbot.com/switch/) {#switch}

Example:

```go
package main

import (
    "fmt"
)

func main() {
    finger := 4
    fmt.Printf("Finger %d is ", finger)
    switch finger {
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 5:
        fmt.Println("Pinky")

    }
}
```

- Duplicate cases are not allowed:
  Duplicate cases with the same constant value are not allowed. If you try to run the program below, the compiler will complain `./prog.go:19:7: duplicate case 4 in switch previous case at ./prog.go:17:7`

```go
package main

import (
    "fmt"
)

func main() {
    finger := 4
    fmt.Printf("Finger %d is ", finger)
    switch finger {
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 4: //duplicate case
        fmt.Println("Another Ring")
    case 5:
        fmt.Println("Pinky")

    }
}
```

- Default case:

```go
package main

import (
    "fmt"
)

func main() {
    switch finger := 8; finger {
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 5:
        fmt.Println("Pinky")
    default: //default case
        fmt.Println("incorrect finger number")
    }
}
```

- Multiple expressions in case
  It is possible to include multiple expressions in a case by separating them with comma.

```go
package main

import (
    "fmt"
)

func main() {
    letter := "i"
    fmt.Printf("Letter %s is a ", letter)
    switch letter {
    case "a", "e", "i", "o", "u": //multiple expressions in case
        fmt.Println("vowel")
    default:
        fmt.Println("not a vowel")
    }
}
```

The above program finds whether letter is a vowel or not. The code case "a", "e", "i", "o", "u": in line no. 11 matches any of the vowels. Since i is a vowel, this program prints
`Letter i is a vowel`

- Expressionless switch
  The expression in a switch is optional and it can be omitted. If the expression is omitted, the switch is considered to be switch true and each of the case expression is evaluated for truth and the corresponding block of code is executed.

```go
package main

import (
    "fmt"
)

func main() {
    num := 75
    switch { // expression is omitted
    case num >= 0 && num <= 50:
        fmt.Printf("%d is greater than 0 and less than 50", num)
    case num >= 51 && num <= 100:
        fmt.Printf("%d is greater than 51 and less than 100", num)
    case num >= 101:
        fmt.Printf("%d is greater than 100", num)
    }

}
```

- Fallthrough:
  In Go, the control comes out of the switch statement immediately after a case is executed. A `fallthrough` statement is used to transfer control to the first statement of the case that is present immediately after the case which has been executed.
  Let's write a program to understand fallthrough. Our program will check whether the input number is less than 50, 100, or 200. For instance, if we input 75, the program will print that 75 is less than both 100 and 200. We will achieve this using `fallthrough`.

```go
package main

import (
    "fmt"
)

func number() int {
        num := 15 * 5
        return num
}

func main() {

    switch num := number(); { //num is not a constant
    case num < 50:
        fmt.Printf("%d is lesser than 50\n", num)
        fallthrough
    case num < 100:
        fmt.Printf("%d is lesser than 100\n", num)
        fallthrough
    case num < 200:
        fmt.Printf("%d is lesser than 200", num)
    }

}
```

- Fallthrough happens even when the case evaluates to false
  There is a subtlety to be considered when using `fallthrough`. Fallthrough will happen even when the case evaluates to false.

```go
package main

import (
    "fmt"
)

func main() {
    switch num := 25; {
    case num < 50:
        fmt.Printf("%d is lesser than 50\n", num)
        fallthrough
    case num > 100:
        fmt.Printf("%d is greater than 100\n", num)
    }

}
```

- Breaking switch:
  The `break` statement can be used to terminate a switch early before it completes. Let's just modify the above example to a contrived one to understand how break works.

```go
package main

import (
    "fmt"
)

func main() {
    switch num := -5; {
    case num < 50:
        if num < 0 {
            break
        }
        fmt.Printf("%d is lesser than 50\n", num)
        fallthrough
    case num < 100:
        fmt.Printf("%d is lesser than 100\n", num)
        fallthrough
    case num < 200:
        fmt.Printf("%d is lesser than 200", num)
    }

}
```

- Breaking the outer for loop
  When the switch case is inside a for loop, there might be a need to terminate the for loop early. This can be done by labeling the for loop and breaking the for loop using that label inside the switch statement. Let's look at an example.
  Let's write a program to generate a random even number.
  We will create an infinite for loop and use a switch case to determine whether the generated random number is even. If it is even, the generated number is printed and the for loop is terminated using its label. The Intn function of the rand package is used to generate non-negative pseudo-random numbers.

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
randloop:
    for {
        switch i := rand.Intn(100); {
        case i%2 == 0:
            fmt.Printf("Generated even number %d", i)
            break randloop
        }
    }

}
```

**Please note that if the break statement is used without the label, the switch statement will only be broken and the loop will continue running. So labeling the loop and using it in the break statement inside the switch is necessary to break the outer for loop.**

## [Arrays and Slices](https://golangbot.com/arrays-and-slices/) {#arrays-and-slices}

#### Arrays

An array is a collection of elements that belong to the same type

- Declaration
  An array belongs to type `[n]T`. `n` denotes the number of elements in an array and `T` represents the type of each element. The number of elements `n` is also a part of the type

```go
package main

import (
    "fmt"
)


func main() {
    var a [3]int //int array with length 3
    fmt.Println(a)
}
```

create the same array using the short hand declaration.

```go
package main

import (
    "fmt"
)

func main() {
    a := [3]int{12, 78, 50} // short hand declaration to create array
    fmt.Println(a)
}
```

ignore the length of the array in the declaration and replace it with `...` and let the compiler find the length for you. This is done in the following program.

```go
package main

import (
    "fmt"
)

func main() {
    a := [...]int{12, 78, 50} // ... makes the compiler determine the length
    fmt.Println(a)
}
```

The size of the array is a part of the type. Hence [5]int and [25]int are distinct types. Because of this, arrays cannot be resized. Don't worry about this restriction since slices exist to overcome this.

```go
package main

func main() {
    a := [3]int{5, 78, 8}
    var b [5]int
    b = a //not possible since [3]int and [5]int are distinct types
}
```

- Arrays are value types
  Arrays in Go are value types and not reference types. This means that when they are assigned to a new variable, a copy of the original array is assigned to the new variable. If changes are made to the new variable, it will not be reflected in the original array.

```go
package main

import "fmt"

func main() {
    a := [...]string{"USA", "China", "India", "Germany", "France"}
    b := a // a copy of a is assigned to b
    b[0] = "Singapore"
    fmt.Println("a is ", a)
    fmt.Println("b is ", b)
}
```

```
a is [USA China India Germany France]
b is [Singapore China India Germany France]
```

- Length of an array: The length of the array is found by passing the array as parameter to the `len` function.

- Iterating arrays using range: The `for` loop can be used to iterate over elements of an array.

```go
package main

import "fmt"

func main() {
    a := [...]float64{67.7, 89.8, 21, 78}
    for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
        fmt.Printf("%d th element of a is %.2f\n", i, a[i])
    }
}
```

Go provides a better and concise way to iterate over an array by using the range form of the for loop. range returns both the index and the value at that index. Let's rewrite the above code using range. We will also find the sum of all elements of the array.

```go
package main

import "fmt"

func main() {
    a := [...]float64{67.7, 89.8, 21, 78}
    sum := float64(0)
    for i, v := range a {//range returns both the index and value
        fmt.Printf("%d the element of a is %.2f\n", i, v)
        sum += v
    }
    fmt.Println("\nsum of all elements of a",sum)
}
```

In case you want only the value and want to ignore the index, you can do this by replacing the index with the \_ blank identifier.

```go
for _, v := range a { //ignores index
}
```

- Multidimensional arrays

```go
package main

import (
    "fmt"
)

func printarray(a [3][2]string) {
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}

func main() {
    a := [3][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
    }
    printarray(a)
    var b [3][2]string
    b[0][0] = "apple"
    b[0][1] = "samsung"
    b[1][0] = "microsoft"
    b[1][1] = "google"
    b[2][0] = "AT&T"
    b[2][1] = "T-Mobile"
    fmt.Printf("\n")
    printarray(b)
}
```

#### Slices

A slice is a convenient, flexible and powerful wrapper on top of an array. Slices do not own any data on their own. They are just references to existing arrays.

- Creating a slice: A slice with elements of type T is represented by []T

```go
package main

import (
    "fmt"
)

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4] //creates a slice from a[1] to a[3]
    fmt.Println(b)
}
```

The syntax `a[start:end]` creates a slice from array a starting from index `start` to index `end - 1`

```go
package main

import (
    "fmt"
)

func main() {
    c := []int{6, 7, 8} //creates and array and returns a slice reference
    fmt.Println(c)
}
```

In the above program in line no. 9, `c := []int{6, 7, 8}` creates an array with 3 integers and returns a slice reference which is stored in c.

- modifying a slice: A slice does not own any data of its own. It is just a representation of the underlying array. Any modifications done to the slice will be reflected in the underlying array

```go
package main

import (
    "fmt"
)

func main() {
    darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
    dslice := darr[2:5]
    fmt.Println("array before",darr)
    for i := range dslice {
        dslice[i]++
    }
    fmt.Println("array after",darr)
}
```

When a number of slices share the same underlying array, the changes that each one makes will be reflected in the array.

```go
package main

import (
    "fmt"
)

func main() {
    numa := [3]int{78, 79 ,80}
    nums1 := numa[:] //creates a slice which contains all elements of the array
    nums2 := numa[:]
    fmt.Println("array before change 1",numa)
    nums1[0] = 100
    fmt.Println("array after modification to slice nums1", numa)
    nums2[1] = 101
    fmt.Println("array after modification to slice nums2", numa)
}
```

In line no. 9, in `numa[:]` the start and end values are missing. The default values for start and end are `0` and `len(numa)` respectively.

- Length and capacity of a slice: The length of the slice is the number of elements in the slice. **The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.**

```go
package main

import (
    "fmt"
)

func main() {
    fruitArray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitSlice := fruitarray[1:3]
    fmt.Printf("length of slice %d capacity %d", len(fruitSlice), cap(fruitSlice)) //length of fruitSlice is 2 and capacity is 6
}
```

In the above program, fruitSlice is created from indexes 1 and 2 of the fruitArray. Hence the length of fruitSlice is 2.

The length of the fruitArray is 7. fruiteSlice is created from index 1 of fruitArray. Hence the capacity of fruitSlice is the no of elements in fruitArray starting from index 1 i.e from orange and that value is 6. Hence the capacity of fruitSlice is 6. The program prints **length of slice 2 capacity 6.**

```go
package main

import (
    "fmt"
)

func main() {
    fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
    fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
    fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity
    fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))
}
```

- Creating a slice using make: `func make([]T, len, cap) []T` can be used to create a slice by passing the type, length and capacity. The capacity parameter is optional and defaults to the length. The make function creates an array and returns a slice reference to it.

```go
package main

import (
    "fmt"
)

func main() {
    i := make([]int, 5, 5)
    fmt.Println(i)
}
```

The values are zeroed by default when a slice is created using make. The above program will output `[0 0 0 0 0]`.

- Appending to a slice: As we already know arrays are restricted to fixed length and their length cannot be increased. Slices are dynamic and new elements can be appended to the slice using append function. The definition of append function is `func append(s []T, x ...T) []T`.

```go
package main

import (
    "fmt"
)

func main() {
    cars := []string{"Ferrari", "Honda", "Ford"}
    fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
    cars = append(cars, "Toyota")
    fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
}
```

It is also possible to append one slice to another using the ... operator.

```go
package main

import (
    "fmt"
)

func main() {
    veggies := []string{"potatoes","tomatoes","brinjal"}
    fruits := []string{"oranges","apples"}
    food := append(veggies, fruits...)
    fmt.Println("food:",food)
}
```

- Passing a slice to a function: Slices can be thought of as being represented internally by a structure type. This is how it looks

```go
type slice struct {
    Length        int
    Capacity      int
    ZerothElement *byte
}
```

A slice contains the length, capacity and a pointer to the zeroth element of the array. When a slice is passed to a function, even though it's passed by value, the pointer variable will refer to the same underlying array. Hence when a slice is passed to a function as parameter, changes made inside the function are visible outside the function too. Lets write a program to check this out.

```go
package main

import (
    "fmt"
)

func subtactOne(numbers []int) {
    for i := range numbers {
        numbers[i] -= 2
    }

}
func main() {
    nos := []int{8, 7, 6}
    fmt.Println("slice before function call", nos)
    subtactOne(nos)                               //function modifies the slice
    fmt.Println("slice after function call", nos) //modifications are visible outside
}
```

output:

```
slice before function call [8 7 6]
slice after function call [6 5 4]
```

- Multidimensional slices:

```go
package main

import (
    "fmt"
)

func main() {
     pls := [][]string {
            {"C", "C++"},
            {"JavaScript"},
            {"Go", "Rust"},
            }
    for _, v1 := range pls {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}
```

output:

```
C C++
JavaScript
Go Rust
```

- Memory Optimization: Slices hold a reference to the underlying array. As long as the slice is in memory, the array cannot be garbage collected. This might be of concern when it comes to memory management. Lets assume that we have a very large array and we are interested in processing only a small part of it. Henceforth we create a slice from that array and start processing the slice. The important thing to be noted here is that the array will still be in memory since the slice references it.

One way to solve this problem is to use the copy function `func copy(dst, src []T) int` to make a copy of that slice. This way we can use the new slice and the original array can be garbage collected.

```go
package main

import (
    "fmt"
)

func countries() []string {
    countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
    neededCountries := countries[:len(countries)-2]
    countriesCpy := make([]string, len(neededCountries))
    copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
    return countriesCpy
}
func main() {
    countriesNeeded := countries()
    fmt.Println(countriesNeeded)
}
```

## [Variadic Functions](https://golangbot.com/variadic-functions/) {#variadic-functions}

Functions in general accept only a fixed number of arguments. A variadic function is a function that accepts a variable number of arguments. If the last parameter of a function definition is prefixed by ellipsis ..., then the function can accept any number of arguments for that parameter. **Only the last parameter of a function can be variadic.**

- Syntax:

```go
func hello(a int, b ...int) {
}
```

In the above function, the parameter b is variadic since it's prefixed by ellipsis and it can accept any number of arguments. This function can be called by using the syntax.

```go
hello(1, 2) //passing one argument "2" to b
hello(5, 6, 7, 8, 9) //passing arguments "6, 7, 8 and 9" to b
```

- Examples and understanding how variadic functions work:

```go
package main

import (
    "fmt"
)

func find(num int, nums ...int) {
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}
func main() {
    find(89, 89, 90, 95)
    find(45, 56, 67, 45, 90, 109)
    find(78, 38, 56, 98)
    find(87)
}
```

The way variadic functions work is by converting the variable number of arguments to a slice of the type of the variadic parameter. For instance, in line no. 22 of the program above, the variable number of arguments to the `find` function are 89, 90, 95. The `find` function expects a variadic `int` argument. Hence these three arguments will be converted by the compiler to a slice of type int `[]int{89, 90, 95}` and then it will be passed to the find function.

```text
type of nums is []int
89 found at index 0 in [89 90 95]

type of nums is []int
45 found at index 2 in [56 67 45 90 109]

type of nums is []int
78 not found in  [38 56 98]

type of nums is []int
87 not found in  []
```

- Slice arguments vs Variadic arguments
