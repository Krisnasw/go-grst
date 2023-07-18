package mysql_model

import (
	"strings"

	"github.com/krisnasw/go-grst/protoc-gen-cdd/descriptor"
	"github.com/iancoleman/strcase"
)

type MysqlModelParam struct {
	*descriptor.MessageDescriptorExt
}

func (mmp MysqlModelParam) GetCrudPackageName() string {
	return strings.Replace(strcase.ToKebab(strings.ToLower("crud-"+mmp.Mysql.TableName)), "-", "_", -1)
}

func (mmp MysqlModelParam) NeedImportTime() bool {
	needImport := false

	if mmp.Mysql != nil {
		needImport = !mmp.Mysql.DisableTimestampTracking || mmp.Mysql.EnableSoftDelete
	}
	if !needImport {
		for _, fext := range mmp.FieldExt {
			if getGoType(fext) == "time.Time" {
				needImport = true
				break
			}
		}
	}
	return needImport
}

func (mmp MysqlModelParam) IsCreatedAt() bool {
	if !mmp.Mysql.DisableTimestampTracking {
		return true
	}
	for _, fext := range mmp.FieldExt {
		if getGoType(fext) == "time.Time" {
			if strings.ToLower(fext.GetName()) == "createdat" || strings.ToLower(fext.GetName()) == "created_at" {
				return true
			}
		}
	}
	return false
}
func (mmp MysqlModelParam) IsUpdatedAt() bool {
	if !mmp.Mysql.DisableTimestampTracking {
		return true
	}
	for _, fext := range mmp.FieldExt {
		if getGoType(fext) == "time.Time" {
			if strings.ToLower(fext.GetName()) == "updatedat" || strings.ToLower(fext.GetName()) == "updated_at" {
				return true
			}
		}
	}
	return false
}

func (mmp MysqlModelParam) GetPrimaryKey() []*descriptor.FieldDescriptorExt {
	fieldpks := []*descriptor.FieldDescriptorExt{}
	for _, f := range mmp.FieldExt {
		if f.MysqlField.PrimaryKey {
			fieldpks = append(fieldpks, f)
		}
	}
	return fieldpks

}

func (mmp MysqlModelParam) GetPrimaryKeyAsString(prefix, suffix, delimiter string, toLower bool, withGoType bool) string {
	fieldpks := mmp.GetPrimaryKey()

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

func (mmp MysqlModelParam) GetPrimaryKeyAsQueryStmt() string {
	fieldpks := mmp.GetPrimaryKey()
	out := []string{}
	for _, pk := range fieldpks {
		out = append(out, pk.MysqlField.ColumnName+" = ?")
	}
	return strings.Join(out, " AND ")
}
