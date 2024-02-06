// Create a race condition by incrementing a value
package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    fmt.Println("Routines:", runtime.NumGoroutine())
    fmt.Println("CPUs:", runtime.NumCPU())

    var wg sync.WaitGroup
    const limit = 100

    var counter = 0

    fmt.Println("Counter:", counter)
    for i := 0; i < limit; i++ {
        wg.Add(1)
        go func(){
            temp := counter
            runtime.Gosched()
            temp++
            counter = temp
            wg.Done()
        }()
        fmt.Println("Counter:", counter)
    }

    wg.Wait()
    fmt.Println("Final Counter:", counter)
}
