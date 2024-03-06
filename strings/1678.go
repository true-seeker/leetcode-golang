package strings

func interpret(command string) string {
	result := ""
	for i := 0; i < len(command); {
		switch command[i] {
		case 'G':
			result += "G"
			i++
		case '(':
			if command[i+1] == ')' {
				result += "o"
				i += 2
			} else {
				result += "al"
				i += 4
			}
		}
	}
	return result
}
