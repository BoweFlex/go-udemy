// Now fix act_3 to not have a race condition
// Use a mutex
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
    var mutagen sync.Mutex

    fmt.Println("Counter:", counter)
    for i := 0; i < limit; i++ {
        wg.Add(1)
        go func(){
            mutagen.Lock()
            temp := counter
            temp++
            counter = temp
            fmt.Println("Counter:", counter)
            mutagen.Unlock()
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("Final Counter:", counter)
}
