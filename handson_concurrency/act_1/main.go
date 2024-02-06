//in addition to the main goroutine, launch two additional goroutines
// each additional goroutine should print something out
// use waitgroups to make sure each goroutine finishes before your program exist
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func main() {
    fmt.Println("Routines:", runtime.NumGoroutine())
    fmt.Println("CPUs:", runtime.NumCPU())

    var wg sync.WaitGroup

    for i := 0; i < runtime.NumCPU(); i++ {
        wg.Add(1)
        go func(){
            time.Sleep(2 * time.Second)
            fmt.Println("Routines:", runtime.NumGoroutine())
            wg.Done()
        }()
    }

    fmt.Println("Routines:", runtime.NumGoroutine())
    wg.Wait()
    fmt.Println("Last Routines:", runtime.NumGoroutine())
}
