// Fix act 3's race condition with an atomic
package main

import (
    "fmt"
    "runtime"
    "sync"
    "sync/atomic"
)

func main() {
    fmt.Println("Routines:", runtime.NumGoroutine())
    fmt.Println("CPUs:", runtime.NumCPU())

    var wg sync.WaitGroup
    const limit = 100

    var incrementer int64

    fmt.Println("Counter:", incrementer)
    for i := 0; i < limit; i++ {
        wg.Add(1)
        go func(){
            atomic.AddInt64(&incrementer, 1)
            fmt.Println("Counter:", atomic.LoadInt64(&incrementer))
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("Final Counter:", incrementer)
}
