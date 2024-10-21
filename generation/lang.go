package generation

var mapping = map[string]string{"java": "java", "python": "Python", "xml": "XML"}

func Map(format string) string {
	f, ok := mapping[format]
	if !ok {
		return ""
	}
	return f
}
