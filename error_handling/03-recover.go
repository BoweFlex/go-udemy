/*How do you recover from a panic?
This example demonstrates the concepts of panic, defer, and recover.
Panic: Tells program to abort while calling any defer/recovers.
Defer: Adds a function to the stack to be called before exiting
Recover: Only really useful inside deferred functions, during normal execution has no effect
but during a panic recover will capture the value and resume normal execution.
*/

package main

import "fmt"

func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}
