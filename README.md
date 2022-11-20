sql语句构建器Quick Start：

1.导入依赖包

	import (
		"database/sql"
		"fmt"
	
		"github.com/CosmicBDry/selfDevelopTools/sqlbuild" ///导入sql语句构建包
	
   		 _ "github.com/go-sql-driver/mysql"
	
	)

2.准备一个sql语句模板

  	var QuerySql = "SELECT `id`, `name`, `province`, `city` FROM `issues`"

3.打开一个db链接池

 	 db, _ := sql.Open("mysql", dsn)

4.将sql模板引入生成一个查询构建器对象

  	queryBuilder :=sqlbuild.NewBuidler(QuerySql)

5.queryBuilder调用Where、OrderBy、Limit等过滤构建器方法，整合成一个完成的sql语句字符串

  	buildSQL :=queryBuilder.Where("id", ">", 10).OrderBy("desc", "id").Limit(10, 30).BuildQuery()

6.将整合的sql语句引入到stmt中执行

  	stmt, _ := db.Prepare(buildSQL)
  
  	rows, err := stmt.Query()
  
  	defer rows.Close()
  
  
7.遍历rows中的查询结果

  	for rows.Next(){
  
      	rows.Scan(&var1,&var2,...)
      
     	fmt.Printf("%v %v...\n",var1,var2,...)
      
     	 ...
      
  	}
  
