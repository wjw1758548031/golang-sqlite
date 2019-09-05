package model

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var Db ModeslDB

type ModeslDB struct {
	Db *sql.DB
}



//操作记得需要先创建表
func init(){
	var err error
	Db.Db, err = sql.Open("sqlite3", "./load.db")
	if err != nil {
		panic(err)
	}
}

func (this *ModeslDB) Create(sqlTable string) {
	//创建表
	res, err := this.Db.Exec(sqlTable)
	if err != nil {
		fmt.Println(err)
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(affect)
}

func (this *ModeslDB) Insert(sql string,data ...interface{}) error {
	//"INSERT INTO userinfo(username, departname, created) values(?,?,?)"
	stmt, err := this.Db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//"wangshubo", "国务院", "2017-04-21"
	_, err = stmt.Exec(data...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	/*affect, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(affect)*/
	return nil
}

func (this *ModeslDB) Delete(sql string,data ...interface{}) error{
	// delete
	//"delete from userinfo where uid=?"
	stmt, err := this.Db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//id
	res, err := stmt.Exec(data...)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}
	//db.Close()
	return nil
}

func (this *ModeslDB) Select(sql string,data ...string){
	//"SELECT * FROM userinfo"
	rows, err := this.Db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	/*var uid int
	var username string
	var department string
	var created time.Time*/
	//result := make([]map[string]interface{},0)
	for i := 0 ; rows.Next() ; i++  {
		//err = rows.Scan(&uid, &username, &department, &created)
		/*for _ , v := range data{
			err = rows.Scan(&result[i]["uid"], &result[i]["uid"], &result[i]["uid"], &result[i]["uid"])

		}*/
		if err != nil {
			fmt.Println(err)
			return
		}
		/*fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)*/
	}
	rows.Close()
}



func (this *ModeslDB) Update(sql string,data ...interface{}) error{
	// update userinfo set username=? where uid=?
	stmt, err := this.Db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = stmt.Exec(data...)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}


