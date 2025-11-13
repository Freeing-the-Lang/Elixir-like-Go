package main

import (
    "fmt"
    "os"
    "Elixir-like-Go/internal/core"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: elixirgo <file.ex>")
        return
    }

    file := os.Args[1]
    err := core.RunFile(file)
    if err != nil {
        fmt.Println("Error:", err)
    }
}
