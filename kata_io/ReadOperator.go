package kata_io
import (
)



func ReadOperator () rune {
	var op rune = 0
	var err error


	for  op <= 0x20  {
		if  op, _, err = reader.ReadRune(); err != nil {
			panic(err)
		}
	}


	return op
}
