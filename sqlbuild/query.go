package sqlbuild

import (
	"fmt"
	"strings"
)

//合成sql查询语句
func (b *Builder) BuildQuery() string {
	if b.base == "" {
		fmt.Println("no sql intput")
		return ""
	}
	return strings.TrimSpace(b.base + " " + strings.Join(b.whereStmt, " ") +
		" " + strings.Join(b.orderStmt, " ") + " " + strings.Join(b.limitStmt, " "))
}
