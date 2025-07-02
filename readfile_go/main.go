package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    data, err := os.ReadFile("test.txt")
    if err != nil {
        log.Fatal("Error reading file:", err)
    }

    fmt.Printf("File content:\n%s\n", string(data))
    fmt.Printf("Number of characters: %d\n", len(data))

    var input string
    fmt.Print("Type something and press enter: ")
    _, err = fmt.Scanln(&input)
    if err != nil {
        fmt.Println("Error reading input:", err)
    } else {
        fmt.Println("You typed:", input)
    }
}
