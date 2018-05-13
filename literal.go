package main

type Literal struct {
    value interface{}
}

func NewLiteral( value interface{}) *Literal {
    return &Literal{
        value,
    }
}
