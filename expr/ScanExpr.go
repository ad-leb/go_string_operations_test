package expr
import (
	"errors"


	"string_calculator/kata_io"
	"string_calculator/strs"
)


func (e *Expr) ScanExpr () {
	this := strs.Read

	e.first = kata_io.ReadQuotedString(MAX_INPUT)
	switch  kata_io.ReadOperator()  {
		case '+':	e.sum()
		case '-':	e.sub()
		case '*':	e.mul()
		case '/':	e.div()
		default	:	panic(errors.New(this + strs.Err_expect_opr))
	}
	if  !kata_io.RestOfLineIsEmpty()  {
		panic(errors.New(this + strs.Err_limit_expr))
	}
}
