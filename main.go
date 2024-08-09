package main
import (
	"fmt"
	"strings"
	"errors"
	"bufio"
	"os"
)



func main () {
	var reader = bufio.NewReader(os.Stdin)

	var one, two string
	var num int
	var op byte
	var res string

	var err error



	for {
		one = read_quoted_string(reader)

		op = 0
		for  op <= 0x20  {
			if op, err = reader.ReadByte(); err != nil {
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
				num = read_number(reader)
				res = strings.Repeat(one, num)
			case '/':	
				num = read_number(reader)
				res = one[:len(one)/num]
			default	:	
				panic(errors.New("unsupported operator has been used."))
		}

		/* There is no anything else in the line */
		if op, err = reader.ReadByte(); err != nil {
			panic(err)
		}
		for  op != 0xa  {
			op, err = reader.ReadByte()
			if  err != nil {
				panic(err)
			}
			if op > 0x20 {
				panic(errors.New("there is too much arguments in expression."))
			}
		}

		
		fmt.Println(res)
		fmt.Println()
	}
}







func read_quoted_string (reader *bufio.Reader) string {
	var c byte = 0
	var tmp string
	var err error


	for  c <= 0x20  {
		if c, err = reader.ReadByte(); err != nil {
			panic(err)
		}
	}
	if c != '"' {
		panic(errors.New("there is a double-quot was expected (begin of string)."))
	}
	if tmp, err = reader.ReadString('"'); err != nil {
		panic(err)
	}
	tmp = strings.Trim(tmp, "\"")


	return tmp
}


func read_number (reader *bufio.Reader) int {
	var number int = 0
	var c byte = 0
	var err error


	for  c <= 0x20  {
		if c, err = reader.ReadByte(); err != nil {
			panic(err)
		}
	}
	for  c > 0x20 {
		if c < '0' || '9' < c {
			panic(errors.New("there is a number was expected."))
		}
		number = (number * 10) + (int(c) & 0xf)
		if c, err = reader.ReadByte(); err != nil {
			panic(err)
		}
	}
	if err = reader.UnreadByte(); err != nil {
		panic(err)
	}


	return number
}
