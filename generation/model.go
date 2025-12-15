package generation

type File struct {
	Lang    string
	Path    string
	Name    string
	Content string
	Prefix  *string
}

type Model []File
