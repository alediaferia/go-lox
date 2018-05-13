package main

type Grouping struct {
	expression *Expression
}

func NewGrouping(expression *Expression) *Grouping {
	return &Grouping{
		expression,
	}
}
