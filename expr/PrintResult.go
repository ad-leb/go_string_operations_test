package expr
import (
	"string_calculator/kata_io"
)

func (e *Expr) PrintResult () {
	kata_io.WriteString(e.result, MAX_OUTPUT)
}
