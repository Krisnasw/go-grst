package mysql_model

import (
	"bytes"
	"go/format"
	"text/template"

	"github.com/krisnasw/go-grst/protoc-gen-cdd/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	tmplRepoMysql = template.Must(template.New("mysql-query").Funcs(template.FuncMap{}).Parse(`
	// Code generated by protoc-gen-cdd. DO NOT EDIT.
	// source: {{.FileExt.GetName}}

	package {{.Mysql.TableName}}
	
	import (
		{{ if .NeedImportTime }} "time" {{ end }}
		"gorm.io/gorm"
	)

	type MysqlDatasource struct {
		db        *gorm.DB
		tableName string
	}

	func NewMysqlDatasource(db *gorm.DB) *MysqlDatasource {
		return &MysqlDatasource{db, "{{.Mysql.TableName}}"}
	}

	func (r *MysqlDatasource) GetByPrimaryKey({{ .GetPrimaryKeyAsString "" "" "," true true }}) (*{{.GetName}}Model, error) {
		result := &{{.GetName}}Model{}
		err := r.db.Table(r.tableName).Where("{{ .GetPrimaryKeyAsQueryStmt }}", {{ .GetPrimaryKeyAsString "" "" "," true false }}).First(&result).Error
		return result, err
	}

	func (r *MysqlDatasource) GetAll() ([]*{{.GetName}}Model, error) {
		result := []*{{.GetName}}Model{}
		err := r.db.Table(r.tableName).Find(&result).Error
		return result, err
	}

	func (r *MysqlDatasource) Create(in {{.GetName}}Model) (*{{.GetName}}Model, error) {
		{{ if or .IsCreatedAt .IsUpdatedAt }} timeNow := time.Now() {{ end }}
		{{ if .IsCreatedAt}} in.CreatedAt = &timeNow {{ end }}
		{{ if .IsUpdatedAt}} in.UpdatedAt = &timeNow {{ end }}

		err := r.db.Table(r.tableName).Create(&in).Error
		if err != nil {
			return nil, err
		}
		return &in, nil
	}

	func (r *MysqlDatasource) Update(in {{.GetName}}Model) (*{{.GetName}}Model, error) {
		{{ if or .IsCreatedAt .IsUpdatedAt }} timeNow := time.Now() {{ end }}
		{{ if .IsCreatedAt}} in.CreatedAt = nil {{ end }}
		{{ if .IsUpdatedAt}} in.UpdatedAt = &timeNow {{ end }}
		err := r.db.Table(r.tableName).Where("{{ .GetPrimaryKeyAsQueryStmt }}", {{ .GetPrimaryKeyAsString "in." "" "," false false }}).Updates(&in).Error
		if err != nil {
			return nil, err
		}
		return &in, nil
	}

	func (r *MysqlDatasource) Delete({{ .GetPrimaryKeyAsString "" "" "," true true }}) error {
		return r.db.Table(r.tableName).Delete(&{{.GetName}}Model{}, "{{ .GetPrimaryKeyAsQueryStmt }}", {{ .GetPrimaryKeyAsString "" "" "," true false }}).Error
	}
	`))
)

func applyTemplateRepoMysql(mmp MysqlModelParam) (*generator.GeneratorResponseFile, error) {
	w := bytes.NewBuffer(nil)
	var tmplData = struct {
		MysqlModelParam
	}{
		mmp,
	}

	if err := tmplRepoMysql.Execute(w, tmplData); err != nil {
		return nil, err
	}

	formatted, err := format.Source([]byte(w.String()))
	if err != nil {
		return nil, err
	}

	return &generator.GeneratorResponseFile{
		Filename:     mmp.Mysql.TableName + "/query.cdd.go",
		Content:      string(formatted),
		GoImportPath: protogen.GoImportPath(""),
	}, nil
}