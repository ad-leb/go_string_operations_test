package kata_io
import (
	"errors"


	"string_calculator/strs"
)



func ReadNumber () int {
	this := strs.ReadNumber
	var number int = 0
	var char rune = 0
	var err error


	for  char <= 0x20  {
		if char, _, err = reader.ReadRune(); err != nil {
			panic(err)
		}
	}
	for  char > 0x20 {
		if char < '0' || '9' < char {
			panic(errors.New(this + strs.Err_expect_num))
		}
		number = (number * 10) + (int(char) & 0xf)
		if char, _, err = reader.ReadRune(); err != nil {
			panic(err)
		}
	}
	if err = reader.UnreadRune(); err != nil {
		panic(err)
	}


	return number
}
