package expr
import (
	"strings"
	"errors"

	"string_calculator/kata_io"
	"string_calculator/strs"
)




func (e *Expr) sum () {
	e.second = kata_io.ReadQuotedString(MAX_INPUT)
	e.result = e.first + e.second
}
func (e *Expr) sub () {
	e.second = kata_io.ReadQuotedString(MAX_INPUT)
	e.result = strings.Replace(e.first, e.second, "", -1)
}
func (e *Expr) mul () {
	this := strs.Op_mul


	if e.number = kata_io.ReadNumber(); e.number > MAX_NUMBER || e.number < MIN_NUMBER {
		panic(errors.New(this + strs.Err_limit_num))
	}
	e.result = strings.Repeat(e.first, e.number)
}
func (e *Expr) div () {
	this := strs.Op_div


	if e.number = kata_io.ReadNumber(); e.number > MAX_NUMBER || e.number < MIN_NUMBER {
		panic(errors.New(this + strs.Err_limit_num))
	}
	e.result = e.first[:len(e.first)/e.number]
}
