package kata_io



func RestOfLineIsEmpty () bool {
	var char rune
	var err error


	if char, _, err = reader.ReadRune(); err != nil {
		panic(err)
	}
	for  char != 0xa  {
		char, _, err = reader.ReadRune()
		if  err != nil {
			panic(err)
		}
		if char > 0x20 {
			return false
		}
	}


	return true
}
