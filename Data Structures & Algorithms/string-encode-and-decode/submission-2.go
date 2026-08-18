type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
		sb.WriteString("№")
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}
	text := []byte("")
	for _, ch := range encoded {
		if ch != '№' {
			text = append(text, byte(ch))
		} else {
			result = append(result, string(text))
			text = []byte("")
		}
	}
	return result
}
