package model

import (
	"fmt"
)

type User struct {
	Uid             int64     `db:"uid" json:"uid"`
	Name            string    `db:"name" json:"name"`
	Ip              string    `db:"ip" json:"ip"`
	IpVisitCount    int64     `db:"ip_visit_count" json:"ip_visit_count"`
	UseIpVisitCount int64     `db:"use_ip_visit_count" json:"use_ip_visit_count"`
	Created         string `db:"created" json:"created"`
}

func (this *ModeslDB) UserSelect() ([]User, error) {
	rows, err := this.Db.Query("SELECT uid,name,ip,ip_visit_count,use_ip_visit_count,created FROM user")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]User, 0)
	for i := 0; rows.Next(); i++ {
		//err = rows.Scan(&uid, &username, &department, &created)
		user := User{}
		err = rows.Scan(&user.Uid, &user.Name, &user.Ip, &user.IpVisitCount, &user.UseIpVisitCount, &user.Created)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		result = append(result, user)
	}
	return result, nil
}


func (this *ModeslDB) UserSelectToParam(param map[string]string) ([]User, error) {
	sql := "SELECT uid,name,ip,ip_visit_count,use_ip_visit_count,created FROM user where 1=1"
	for k,v := range param{
		sql += " and " + k + " = " + "'"+v+"'"
	}
	fmt.Println(sql)
	rows, err := this.Db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]User, 0)
	for i := 0; rows.Next(); i++ {
		//err = rows.Scan(&uid, &username, &department, &created)
		user := User{}
		err = rows.Scan(&user.Uid, &user.Name, &user.Ip, &user.IpVisitCount, &user.UseIpVisitCount, &user.Created)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		result = append(result, user)
	}
	return result, nil
}
