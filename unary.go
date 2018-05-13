package main

type Unary struct {
    operator Token
    right *Expression
}

func NewUnary( operator Token, right *Expression) *Unary {
    return &Unary{
        operator,
        right,
    }
}
