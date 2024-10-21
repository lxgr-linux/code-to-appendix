package generation

type File struct {
	Lang    string
	Path    string
	Name    string
	Content string
}

type Model []File
