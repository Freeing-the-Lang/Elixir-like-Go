package core

import (
    "fmt"
    "os"
)

func RunFile(path string) error {
    src, err := os.ReadFile(path)
    if err != nil {
        return err
    }

    fmt.Println(">>> Running Elixir-like-Go program")
    fmt.Println(string(src))

    // TODO:
    // 1. Lexer → tokens
    // 2. Parser → AST
    // 3. Evaluator → runtime 실행
    return nil
}
