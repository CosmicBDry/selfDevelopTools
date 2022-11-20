package sqlbuild

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Builder struct {
	base          string
	whereStmt     []string
	limitStmt     []string
	orderbyMethod string
	orderStmt     []string
}
//newy一个构建器对象
func NewBuidler(sqlQuery string) *Builder {
	return &Builder{
		base:          sqlQuery,
		orderbyMethod: "desc",
	}
}
//where语句构建器
func (b *Builder) Where(ColumnName, OperatorMethod string, Value interface{}) *Builder {
	ColumnName = strings.TrimSpace(ColumnName)
	OperatorMethod = strings.TrimSpace(OperatorMethod)
	if len(ColumnName) != 0 && len(OperatorMethod) != 0 && Value != nil {
		dataType := reflect.TypeOf(Value).String()
		if dataType == "string" {
			values := strings.TrimSpace(Value.(string))
			if len(values) > 0 {
				b.whereStmt = append(b.whereStmt, "WHERE", ColumnName, OperatorMethod, `"`+values+`"`)
			}
		} else if dataType == "int" {
			b.whereStmt = append(b.whereStmt, "WHERE", ColumnName, OperatorMethod, strconv.Itoa(Value.(int)))
		} else if dataType == "int64" {
			b.whereStmt = append(b.whereStmt, "WHERE", ColumnName, OperatorMethod, strconv.FormatInt(Value.(int64), 10))
		} else {
			fmt.Println("datatype error,value  is only : `string,int,int64`")
		}
	}
	return b
}
//orderby排序构建器
func (b *Builder) OrderBy(Method string, ColumnName ...string) *Builder {
	if strings.TrimSpace(Method) != "" {
		b.orderbyMethod = Method
	}
	if len(ColumnName) > 0 {
		b.orderStmt = append(b.orderStmt, "ORDER BY", strings.Join(ColumnName, ","), b.orderbyMethod)
	}
	return b
}
//limit、offset翻页构建器
func (b *Builder) Limit(Offset, Limit uint64) *Builder {
	if Limit > 0 {
		OffsetStr := strconv.FormatUint(Offset, 10)
		limitStr := strconv.FormatUint(Limit, 10)
		b.limitStmt = append(b.limitStmt, "limit", OffsetStr+","+limitStr)
	}
	return b
}
//将以上构建器对象合成sql查询语句
func (b *Builder) BuildQuery() string {
	if b.base == "" {
		fmt.Println("no sql intput")
		return ""
	}
	return strings.TrimSpace(b.base + " " + strings.Join(b.whereStmt, " ") +
		" " + strings.Join(b.orderStmt, " ") + " " + strings.Join(b.limitStmt, " "))
}
