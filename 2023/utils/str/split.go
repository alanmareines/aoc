package str

func Split(input string, delimiter string) []string {
	var result []string
	var ref int
	for i, c := range input {
		if string(c) == delimiter {
			result = append(result, input[ref:i])
			ref = i
		}
		if i == len(input)-1 {
			result = append(result, input[ref:])
		}
	}
	return result
}
