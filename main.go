package main

import (
	"log"
	"os"

	"github.com/lxgr-linux/code-to-appendix/config"
	"github.com/lxgr-linux/code-to-appendix/generation"
)

func main() {
	cfg := getConfig()

	names, err := getFiles(".")
	if err != nil {
		log.Panicf("Failed reading files: '%s'\n", err)
	}

	outp, err := generation.Generate(cfg, names)
	if err != nil {
		log.Panicf("Failed generating files: '%s'\n", err)
	}

	err = os.WriteFile(cfg.OutputFile, outp, os.ModePerm)
	if err != nil {
		log.Panicf("Failed to write file: '%s'\n", err)
	}
}

func getConfig() (cfg *config.Config) {
	if len(os.Args) >= 2 {
		cfg, err := config.Load(os.Args[1])
		if err != nil {
			log.Panicf("Couldn't load condig: '%s'\n", err)
		}
		return cfg
	}
	return config.Default()
}

func getFiles(basePath string) (names []string, err error) {
	dir, err := os.ReadDir(basePath)
	if err != nil {
		return []string{}, err
	}

	for _, file := range dir {
		name := basePath + "/" + file.Name()
		if file.IsDir() {
			subNames, err := getFiles(name)
			if err != nil {
				return []string{}, err
			}
			names = append(names, subNames...)
		} else {
			names = append(names, name)
		}
	}

	return
}
