package generation

var mapping = map[string]string{"java": "java", "py": "Python", "xml": "XML", "sh": "bash"}

func Map(format string) string {
	f, ok := mapping[format]
	if !ok {
		return ""
	}
	return f
}
