package generation

import (
	"bytes"
	_ "embed"
	"os"
	"regexp"
	"slices"
	"strings"
	"text/template"

	"github.com/lxgr-linux/code-to-appendix/config"
)

//go:embed appendix.pb.tex.tmpl
var appendixTmpl string

func Generate(cfg *config.Config, fileNames []string) ([]byte, error) {
	var model Model

filesLoop:
	for _, name := range fileNames {
		format := getFormat(name)

		for _, exp := range cfg.ExcludeFiles {
			match, err := regexp.Match(exp, []byte(name))
			if err != nil {
				return []byte{}, err
			}
			if match {
				continue filesLoop
			}
		}

		if slices.Contains(cfg.ExcludeFormats, format) {
			continue
		}

		content, err := os.ReadFile(name)
		if err != nil {
			return []byte{}, err
		}

		splidPath := strings.Split(name, "/")

		model = append(model, File{
			Lang:    Map(format),
			Path:    cleanPath(name),
			Name:    cleanTitle(splidPath[len(splidPath)-1]),
			Content: string(content),
			Prefix:  cfg.Prefix,
		})
	}

	tmpl, err := template.New("appendix").Parse(appendixTmpl)
	if err != nil {
		return []byte{}, err
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, model)
	if err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), err
}

func getFormat(name string) string {
	splid := strings.Split(name, ".")
	if len(splid) == 1 {
		return ""
	}
	return splid[len(splid)-1]
}

func cleanPath(path string) string {
	return cleanName(strings.TrimPrefix(path, "./"))
}

func cleanName(name string) string {
	return strings.ReplaceAll(name, "_", "\\string_")
}

func cleanTitle(name string) string {
	return strings.ReplaceAll(name, "_", "\\textunderscore{}")
}
