package jsonparser

func ParseJSONSeq(input string) []string {
	stack := make([]string, 0)
	res := make([]string, 0)
	inQ := false
	start := 0
	for i, k := range input {
		if k == '"' {
			if input[i-1] != '\\' ||
				(input[i-1] == '\\' && input[i-2] == '\\') {
				inQ = !inQ
			}
		}

		if inQ {
			continue
		}

		if k == '{' {
			stack = append(stack, "{")
		}

		if k == '}' {
			res = append(res, input[start:i+1])
			start = i + 1
			stack = stack[:len(stack)-1]
		}
	}

	return res
}
