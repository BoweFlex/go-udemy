// You can also solve blocking with buffer channels.
// This has its uses, but is not recommended if you can think through and plan for proper passing.
package main

import "fmt"

func main() {
    c := make(chan int, 1) // Buffer channel, allows 1 value to sit at a time

    c <- 42 
    c <- 43 // However this fails because only one value is allowed and we haven't pulled anything off. 

    fmt.Println(<-c)
}
