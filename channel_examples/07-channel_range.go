//If you're like me, at this point you're asking:
// "Why would I want a single direction channel if I can only send values I can't get back out, or vice versa?"
// Here's a simple example of how they would actually be used. (hint: it has to do with routines)

package main

import "fmt"

func main() {
    c := make(chan int)

    go func() {
        for i := 1; i <= 5; i++ {
            c <- i
        }
        close(c) // This closes the original channel, even if it is in a separate function.
    }()

    for v := range c {
        // Ranging over channel will deadlock if it is not closed
        fmt.Println(v)
    }

    fmt.Println("about to exit")
}
