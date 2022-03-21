package shell

func IsSpace(data string) bool {
	return data == " "
}

func IsOperator(data string) bool {
	if _, found := OP[data]; found {
		return true
	}

	return false
}
