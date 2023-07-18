package gengo_mysql_query

import (
	"bufio"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
)

type generatorResponseFile struct {
	outputPath string
	content    string
}

func Generate(inputFile, modelName string, outputDir string) error {
	packageName := getPackageNameFromFile(inputFile)
	mysqlModel := generateMysqlModelsFromFile(inputFile, modelName)
	if mysqlModel != nil {
		fMysqlQuery, err := applyTemplateQuery(*mysqlModel, packageName, outputDir, inputFile)
		if err != nil {
			return err
		}

		files := []*generatorResponseFile{}
		files = append(files, fMysqlQuery)

		for _, f := range files {
			err := f.generateFile()
			if err != nil {
				return err
			}
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

func generateMysqlModelsFromFile(inputFile string, structName string) *MysqlModel {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	src, err := ioutil.ReadAll(file)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	for _, decl := range f.Decls {
		switch declType := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range declType.Specs {
				switch typeSpec := spec.(type) {
				case *ast.TypeSpec:
					gen := MysqlModel{}
					gen.Name = typeSpec.Name.Name
					gen.TableName = strcase.ToSnake(gen.Name)
					gen.PrimaryKeys = []GoField{}
					if strings.ToLower(structName) != strings.ToLower(gen.Name) {
						continue
					}

					switch typeSpec.Type.(type) {
					case *ast.StructType:
						structType := typeSpec.Type.(*ast.StructType)
						for _, field := range structType.Fields.List {
							if field.Doc != nil {
								for _, comment := range field.Doc.List {
									if strings.Contains(comment.Text, "table_name") && strings.Contains(comment.Text, "=") {
										trimmedComment := strings.ReplaceAll(strings.ReplaceAll(comment.Text, "table_name", ""), "=", "")
										trimmedComment = strings.ReplaceAll(trimmedComment, "//", "")
										trimmedComment = strings.ReplaceAll(trimmedComment, `"`, "")
										trimmedComment = strings.TrimSpace(trimmedComment)
										if trimmedComment != "" {
											gen.TableName = strcase.ToSnake(trimmedComment)
										}
									}
								}
							}

							typeExpr := field.Type
							start := typeExpr.Pos() - 1
							end := typeExpr.End() - 1
							typeInSource := src[start:end]

							tags := ""
							if field.Tag != nil {
								tags = field.Tag.Value
							}

							if strings.Contains(strings.ToLower(tags), "primary_key") {
								gen.PrimaryKeys = append(gen.PrimaryKeys, GoField{
									Name: field.Names[0].Name,
									Type: string(typeInSource),
									Tag:  tags,
								})
							} else if strings.ToLower(field.Names[0].Name) == "createdat" {
								t := NewMysqlTimestampTrackerType(string(typeInSource))
								if t != MysqlTimestampTracker_Unknown {
									gen.IsCreatedAt = true
									gen.CreatedAtType = t

								}
							} else if strings.ToLower(field.Names[0].Name) == "updatedat" {
								t := NewMysqlTimestampTrackerType(string(typeInSource))
								if t != MysqlTimestampTracker_Unknown {
									gen.IsUpdatedAt = true
									gen.UpdatedAtType = t
								}
							}
						}

					}
					return &gen
				}
			}
		}
	}
	return nil
}
func getPackageNameFromFile(inputFile string) string {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	src, err := ioutil.ReadAll(file)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	return f.Name.Name
}
