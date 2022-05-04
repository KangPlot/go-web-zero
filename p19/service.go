package main

import (
	"database/sql"
	"log"
)

func getOne(id int) (a app, err error) {
	a = app{}

	// [order]  是sql里面的关键字，加上【】
	err = db.QueryRow("select id, name, status,"+
		"level, [order] from dbo.App").Scan(
		&a.ID, &a.name, &a.status, &a.level, &a.order)
	return
}

// 返回的切片【】
func getMany(id int) (apps []app, err error) {

	// [order]  是sql里面的关键字，加上【】
	rows, err := db.Query("select id, name, status,"+
		"level, [order] from dbo.App where id > @id", sql.Named("Id", id))

	//对rows 遍历

	for rows.Next() {
		a := app{}
		err = rows.Scan(
			&a.ID, &a.name, &a.status, &a.level, &a.order)
		if err != nil {
			log.Fatalln(err.Error())
		}
		apps = append(apps, a)
	}
	return apps, err

}
