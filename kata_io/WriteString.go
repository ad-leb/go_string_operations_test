package kata_io



func WriteString (str string, max_output int) {
	var char rune
	var count int
	var err error

	

	/* Count characters (not bytes) while puting them */
	count = 0
	for _, char = range str {
		count++
		if count > max_output {
			if _, err = writer.WriteString("..."); err != nil {
				panic(err)
			}
			break
		}
		_, err = writer.WriteRune(char)
		if  err != nil  {
			panic(err)
		}
	}

	/* Write 2 newline-characters */
	if _, err = writer.WriteString("\n\n"); err != nil {
		panic(err)
	}




	/* Put all of this into the Stdout */
	if err = writer.Flush(); err != nil {
		panic(err)
	}
}
