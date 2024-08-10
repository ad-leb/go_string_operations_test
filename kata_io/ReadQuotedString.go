package kata_io
import (
	"strings"
	"errors"


	"string_calculator/strs"
)



func ReadQuotedString (max_input int) string {
	this := strs.ReadQuotedString

	var tmp strings.Builder = strings.Builder{}
	var char rune = 0
	var count int = 0
	var err error


	for  char <= 0x20  {
		if char, _, err = reader.ReadRune(); err != nil {
			panic(err)
		}
	}
	if char != '"' {
		panic(errors.New(this + strs.Err_expect_str))
	}


	if char, _, err = reader.ReadRune(); err != nil {
		panic(err)
	}
	for char != '"' {
		if count > max_input {
			panic(errors.New(this + strs.Err_limit_str))
		} else {
			tmp.WriteRune(char)
			count++
		}

		if char, _, err = reader.ReadRune(); err != nil {
			panic(err)
		}
	}



	return tmp.String()
}
