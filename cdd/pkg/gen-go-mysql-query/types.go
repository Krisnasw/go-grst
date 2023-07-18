package gengo_mysql_query

import (
	"strings"

	"github.com/iancoleman/strcase"
)

type GoField struct {
	Name string
	Type string
	Tag  string
}

type MysqlTimestampTrackerType string

const (
	MysqlTimestampTracker_Unknown = ""
	MysqlTimestampTracker_Time    = "time.Time"
	MysqlTimestampTracker_TimePtr = "*time.Time"
	MysqlTimestampTracker_Int     = "int"
	MysqlTimestampTracker_IntPtr  = "*int"
)

func NewMysqlTimestampTrackerType(in string) MysqlTimestampTrackerType {
	switch in {
	case "time.Time":
		return MysqlTimestampTracker_Time
	case "*time.Time":
		return MysqlTimestampTracker_TimePtr
	case "int":
		return MysqlTimestampTracker_Int
	case "*int":
		return MysqlTimestampTracker_IntPtr
	}
	return MysqlTimestampTracker_Unknown
}

type MysqlModel struct {
	Name          string
	TableName     string
	PrimaryKeys   []GoField
	IsCreatedAt   bool
	CreatedAtType MysqlTimestampTrackerType
	IsUpdatedAt   bool
	UpdatedAtType MysqlTimestampTrackerType
}

func (model MysqlModel) GetName() string {
	return model.Name
}
func (model MysqlModel) NeedImportTime() bool {
	return (model.IsCreatedAt && (model.CreatedAtType == MysqlTimestampTracker_Time || model.CreatedAtType == MysqlTimestampTracker_TimePtr)) ||
		(model.IsUpdatedAt && (model.UpdatedAtType == MysqlTimestampTracker_Time || model.UpdatedAtType == MysqlTimestampTracker_TimePtr))

}

func (model MysqlModel) HasPrimaryKey() bool {
	return len(model.PrimaryKeys) > 0
}
func (model MysqlModel) GetPrimaryKeyAsString(prefix, suffix, delimiter string, toLower bool, withGoType bool) string {
	fieldpks := model.PrimaryKeys

	out := []string{}
	for _, pk := range fieldpks {
		pkName := pk.Name
		if toLower {
			pkName = strings.ToLower(pkName)
		}
		tmpOut := prefix + pkName
		if withGoType {
			tmpOut += " " + pk.Type
		}
		tmpOut += suffix
		out = append(out, tmpOut)
	}
	return strings.Join(out, delimiter)
}

func (model MysqlModel) GetPrimaryKeyAsQueryStmt() string {
	fieldpks := model.PrimaryKeys
	out := []string{}

	for _, pk := range fieldpks {
		columnName := getStringInBetween(pk.Tag, "column:", ";")
		if columnName == "" {
			columnName = getStringInBetween(pk.Tag, "column:", `"`)
		}
		if columnName == "" {
			columnName = getStringInBetween(pk.Tag, "column:", `,`)
		}
		if columnName == "" {
			columnName = strcase.ToSnake(pk.Name)
		}
		out = append(out, columnName+" = ?")
	}
	return strings.Join(out, " AND ")
}
func getStringInBetween(str string, patternStart string, patternEnd string) (result string) {
	s := strings.Index(str, patternStart)
	if s == -1 {
		return
	}
	s += len(patternStart)
	e := strings.Index(str[s:], patternEnd)
	if e == -1 {
		return
	}
	return str[s : s+e]
}
