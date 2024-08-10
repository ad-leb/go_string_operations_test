package main
import (
	"string_calculator/expr"
)




func main () {
	var expr *expr.Expr = &expr.Expr{}


	for {
		expr.ScanExpr()
		expr.PrintResult()
	}

}
