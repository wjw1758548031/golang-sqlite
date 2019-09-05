package main

import (
	"database/sql"
	"fmt"
	models "test/test1/model"
)

//用命令行运行
func main(){
	var err error
	models.Db.Db, err = sql.Open("sqlite3", "./load.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlTable := `
    CREATE TABLE IF NOT EXISTS user(
        uid  INTEGER PRIMARY KEY AUTOINCREMENT,
        name VARCHAR(64) NULL,
        ip   VARCHAR(64) NULL,
		ip_visit_count INT NULL,
		use_ip_visit_count INT NULL,
        created DATETIME NULL，
    );
    `
	models.Db.Create(sqlTable)
}


