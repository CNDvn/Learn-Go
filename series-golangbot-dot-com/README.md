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
[Switch Statement](https://github.com/CNDvn/Learn-Go/tree/main/series-golangbot-dot-com#switch)

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

## [Switch Statement](https://golangbot.com/switch/) {https://github.com/CNDvn/Learn-Go/tree/main/series-golangbot-dot-com#switch}

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
