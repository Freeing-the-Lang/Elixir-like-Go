package lexer

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

func Lex(input string) []Token {
    tokens := []Token{}
    // TODO: 간단한 Elixir 토큰화 규칙 추가
    return tokens
}
