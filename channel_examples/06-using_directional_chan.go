//If you're like me, at this point you're asking:
// "Why would I want a single direction channel if I can only send values I can't get back out, or vice versa?"
// Here's a simple example of how they would actually be used. (hint: it has to do with routines)

package main

import "fmt"

func main() {
    c := make(chan int)

    // This is going to do something and send a value to the channel
    go foo(c) // Launch in a routine

    //This is going to receive from the channel and do something with it.
    //Without a wait group, both routines would kick off and main would report success before 
    //bar could read from the channel.    
    //go bar(c)
    // You can either implement a wait group, or:
    bar(c) // not launch a routine, and this will hang until there's a value in the channel

    fmt.Println("about to exit")
}

func foo(c chan<- int){
    c <- 42
}

func bar(c <-chan int){
    fmt.Println(<-c)
}
