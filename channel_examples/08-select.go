// Like a switch statement, but for channels.

package main

import "fmt"

func main() {
    even := make(chan int)
    odd := make(chan int)
    quit := make(chan int)

    // Send to channels
    go send(even, odd, quit)

    receive(even, odd, quit)

    fmt.Println("about to exit")
}

func send(e, o, q chan<- int) {
    for i := 1; i <= 100; i++ {
        if i % 2 == 0 {
            e <- i
        } else {
            o <- i
        }
    }
    // close(e)
    // close(o) // These cause some weird behavior at the end, not sure why
    // Maybe they get closed and receiver can pull some zeros before quit is filled and pulled from?
    q <- 0
}

func receive(e, o, q <-chan int) {
    for { // infinite loop to make sure I receive all values, needs to have a case to return
        select { // Pulls from whatever channel is available
        case v := <- e: // If able to pull from even
            fmt.Println("from the even channel:", v)
        case v := <- o: // If able to pull from odd
            fmt.Println("from the odd channel:", v)
        case v := <- q: // If able to pull from quit
            fmt.Println("from the quit channel:", v)
            return
        }
    }
}

