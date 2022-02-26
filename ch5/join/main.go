package main

func Join(sep string, elements ...string) string {
	result := ""
	for _, s := range elements {
		if len(result) > 0 {
			result += sep + s
		} else {
			result += s
		}
	}
	return result
}
