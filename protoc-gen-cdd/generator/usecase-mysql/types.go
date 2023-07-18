package usecase_mysql

import (
	"strings"

	"github.com/krisnasw/go-grst/protoc-gen-cdd/descriptor"
	"github.com/iancoleman/strcase"
)

type ScaffoldMysql struct {
	*descriptor.MessageDescriptorExt
	GoModuleName string
}

func (s ScaffoldMysql) GetCrudPackageName() string {
	return strings.Replace(strcase.ToKebab(strings.ToLower("crud-"+s.Mysql.TableName)), "-", "_", -1)
}

func (s ScaffoldMysql) NeedImportTime() bool {
	needImport := false

	if s.Mysql != nil {
		needImport = !s.Mysql.DisableTimestampTracking || s.Mysql.EnableSoftDelete
	}
	if !needImport {
		for _, fext := range s.FieldExt {
			if getGoType(fext) == "time.Time" {
				needImport = true
				break
			}
		}
	}
	return needImport
}

func (s ScaffoldMysql) GetPrimaryKey() []*descriptor.FieldDescriptorExt {
	fieldpks := []*descriptor.FieldDescriptorExt{}
	for _, f := range s.FieldExt {
		if f.MysqlField.PrimaryKey {
			fieldpks = append(fieldpks, f)
		}
	}
	return fieldpks

}

func (s ScaffoldMysql) GetPrimaryKeyAsString(prefix, suffix, delimiter string, toLower bool, withGoType bool) string {
	fieldpks := s.GetPrimaryKey()

	out := []string{}
	for _, pk := range fieldpks {
		pkName := pk.GetName()
		if toLower {
			pkName = strings.ToLower(pkName)
		}
		tmpOut := prefix + pkName
		if withGoType {
			tmpOut += " " + getGoStandartType(pk)
		}
		tmpOut += suffix
		out = append(out, tmpOut)
	}
	return strings.Join(out, delimiter)
}

func (s ScaffoldMysql) GetPrimaryKeyAsQueryStmt() string {
	fieldpks := s.GetPrimaryKey()
	out := []string{}
	for _, pk := range fieldpks {
		out = append(out, pk.MysqlField.ColumnName+" = ?")
	}
	return strings.Join(out, " AND ")
}
