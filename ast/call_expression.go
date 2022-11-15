package ast

type CallExpression struct {
	Function   Identifier
	Parameters []Identifier
}

func (c CallExpression) TokenLiteral() string {
	//TODO implement me
	panic("implement me")
}

func (c CallExpression) String() string {
	//TODO implement me
	panic("implement me")
}

func (c CallExpression) expression() {
	//TODO implement me
	panic("implement me")
}
