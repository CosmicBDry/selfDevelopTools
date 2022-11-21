package sqlbuild

import (
	"fmt"
	"strings"
)

//合成sql查询语句
func (b *Builder) BuildQuery() string {
	if b.base == "" {
		fmt.Println("no sqlTemplate input")
		return ""
	}
	b.joinStmt = strings.TrimSpace(b.base + " " + strings.Join(b.whereStmt, " ") +
		" " + strings.Join(b.orderStmt, " ") + " " + strings.Join(b.limitStmt, " "))
	return b.joinStmt
}
