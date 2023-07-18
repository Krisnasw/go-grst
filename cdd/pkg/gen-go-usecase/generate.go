package gengousecase

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type generatorResponseFile struct {
	outputPath string
	content    string
}

func Generate(usecaseName string, dir string) error {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	usecaseName = strings.TrimSpace(usecaseName)
	usecaseName = reg.ReplaceAllString(usecaseName, "_")
	if dir == "" {
		dir = "./"
	}
	folderpath := dir + "/" + usecaseName
	os.MkdirAll(folderpath, os.ModePerm)

	files := []*generatorResponseFile{}
	fUsecaseError, err := applyTemplateUseCaseErrors(usecaseName, folderpath)
	if err != nil {
		return err
	}
	files = append(files, fUsecaseError)

	fUsecaseImpl, err := applyTemplateUseCaseImpl(usecaseName, folderpath)
	if err != nil {
		return err
	}
	files = append(files, fUsecaseImpl)

	fUsecaseIntf, err := applyTemplateUseCaseIntf(usecaseName, folderpath)
	if err != nil {
		return err
	}
	files = append(files, fUsecaseIntf)

	// fUsecaseRepoImpl, err := applyTemplateUseCaseRepoImpl(usecaseName, folderpath)
	// if err != nil {
	// 	return err
	// }
	// files = append(files, fUsecaseRepoImpl)

	for _, f := range files {
		err := f.generateFile()
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *generatorResponseFile) generateFile() error {
	f, err := os.Create(g.outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(g.content)
	w.Flush()
	return nil
}
