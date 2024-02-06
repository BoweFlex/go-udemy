// You can specify a channel in three ways:
// 1. bidirectional (or general)
// or specific directions:
// 2. send only
// 3. receive only
// These can be read in order to make sure you understand, i.e:
// <-chan int "receive from channel integers"
// chan<- int "channel you send integers to"
package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) //receive
	cs := make(chan<- int) //send

    //You can't assign a single direction channel to a bidirectional channel
    //You would lose information, specific -> general
    //c = cr
    //c = cs

    //But general to specific works since you are only adding information
    cr = c
    cs = c

    //Same for temporary conversions
    fmt.Printf("c\t%T\n", (chan<- int)(c)) // this works
	//fmt.Printf("c\t%T\n", (chan int)(cr)) // this fails

    cs <- 42 // This works because it is a channel you send to
    //cr <- 42 // This doesn't work because it's a channel you receive from

    //fmt.Println(<-cs) // This fails because it's a send channel
    fmt.Println(<-cr) // This works because it's a receive channel
}
