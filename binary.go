package main

type Binary struct {
    left *Expression
    operator Token
    right *Expression
}

func NewBinary( left *Expression, operator Token, right *Expression) *Binary {
    return &Binary{
        left,
        operator,
        right,
    }
}
