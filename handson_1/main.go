package main

import "fmt"

func main() {
    fmt.Println("This is some text with emojis: 🥳🤯🫥\n")
    fmt.Println(`This is a 
    raw string
    that retains
    new lines and odd characters.
    `)
    fmt.Println("This just says hi\n")
    fmt.Printf("This is a line with a %s.\n", "string being formatted in")
}
