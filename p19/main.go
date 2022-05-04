package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	// 数据库驱动， _ 自我注册init 会自动调用。把包变成_ , 不会直接使用。
	//没有变量，不影响代码里面的逻辑
	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB // 指向数据库
func main() {
	// 连接字符串
	connStr := fmt.Sprintf("server=%s;user id = %s; "+
		""+
		"password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	fmt.Println(connStr)

	var err error
	// 这个db 需要是全局的，不能用:= 
	db, err = sql.Open("sqlserver", connStr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	ctx := context.Background()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// 测试是否连接数据库
	fmt.Println("connected!")

	// 查询
	one, err := getOne(103)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(one)
}

const (
	server   = "xxxx.xxx"
	port     = 1433
	user     = "xxx"
	password = "123"
	database = "go-db"
)
