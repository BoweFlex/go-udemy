//Comma, ok can be applied to channels as well
package main

import "fmt"

func main() {
    c := make(chan int)

    go func() {
        c <- 42
        close(c)
    }()

    v, ok := <-c

    fmt.Println(v, ok)

    v, ok = <-c // c no longer open to be pulled from

    fmt.Println(v, ok)
}
