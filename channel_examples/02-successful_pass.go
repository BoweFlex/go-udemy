// You can solve the blocking issue by "passing the baton" from one routine to another.
package main

import "fmt"

func main() {
    c := make(chan int)

    go func(){
        c <- 42 // this still blocks, but is blocking secondary routine so main can continue
    }()

    fmt.Println(<-c) // This takes over the channel, then secondary routine ends
}
