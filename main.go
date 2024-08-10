package main
import (
	"strings"
	"errors"
	"bufio"
	"os"
)


const MAX_NUMBER = 10
const MIN_NUMBER = 1
const MAX_INPUT = 10
const MAX_OUTPUT = 40



func main () {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)

	var one, two string
	var num int
	var op rune
	var res string

	var runes int
	var err error




	/* Main cycle */
	for {
		/* Read expression */
		one = read_quoted_string(reader)

		op = 0
		for  op <= 0x20  {
			if op, _, err = reader.ReadRune(); err != nil {
				panic(err)
			}
		}
		switch op {
			case '+':	
				two = read_quoted_string(reader)
				res = one + two

			case '-':	
				two = read_quoted_string(reader)
				res = strings.Replace(one, two, "", -1)

			case '*':	
				if num = read_number(reader); num > MAX_NUMBER || num < MIN_NUMBER {
					panic(errors.New("can't use number bigger than 10 and less than 1"))
				}
				res = strings.Repeat(one, num)
			case '/':	
				if num = read_number(reader); num > MAX_NUMBER || num < MIN_NUMBER {
					panic(errors.New("can't use number bigger than 10 and less than 1"))
				}
				res = one[:len(one)/num]
			default	:	
				panic(errors.New("unsupported operator has been used."))
		}

		/* There is no anything else in the line */
		if op, _, err = reader.ReadRune(); err != nil {
			panic(err)
		}
		for  op != 0xa  {
			op, _, err = reader.ReadRune()
			if  err != nil {
				panic(err)
			}
			if op > 0x20 {
				panic(errors.New("too much arguments in expression."))
			}
		}





		/* Output string */
		runes = 0
		for _, char := range res {
			runes++
			if runes > MAX_OUTPUT {
				writer.WriteString("...")
				break
			}
			writer.WriteRune(char)
		}
		writer.WriteByte('\n')
		writer.WriteByte('\n')
		writer.Flush()
	}
	/* End of main cycle */

}










func read_quoted_string (reader *bufio.Reader) string {
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
		panic(errors.New("there is a double-quot was expected (begin of string)."))
	}


	if char, _, err = reader.ReadRune(); err != nil {
		panic(err)
	}
	for char != '"' {
		if count > MAX_INPUT {
			panic(errors.New("only up to 10 character strings is required."))
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


func read_number (reader *bufio.Reader) int {
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
			panic(errors.New("there is a number was expected."))
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
