package generation

var mapping = map[string]string{
	"java": "java", "py": "Python", "xml": "XML", "sh": "bash",
	"go": "go", "js": "JavaScript", "ts": "JavaScript", "kt": "Kotlin", "kts": "Kotlin", "rs": "Rust",
}

func Map(format string) string {
	f, ok := mapping[format]
	if !ok {
		return ""
	}
	return f
}
