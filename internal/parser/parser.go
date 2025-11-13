package parser

import "Elixir-like-Go/internal/lexer"

type Node interface{}

func Parse(tokens []lexer.Token) Node {
    // TODO: Elixir의 패턴매칭 문법을 Go AST로 변환
    return nil
}
