package model

import (
	"fmt"
	"ssevven/utils"
)

type User struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//添加用户方法1//可以防止Sql注入?
func (user *User) AddUser() error {
	//sql 语句 //注册
	sqlStr := "insert into user(username,password) values(?,?)"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译处理异常=", err)
		return err
	}
	//执行
	_, err2 := inStmt.Exec(user.Username, user.Password)
	if err2 != nil {
		fmt.Println("执行出现异常=", err2)
		return err2
	}
	return nil
}
func (user *User) GetUserByUsername() (*User, error) {
	sqlStr := "select uid, username, password from user where username = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, user.Username)
	//声明变量
	var (
		uid      int
		username string
		password string
	)
	err := row.Scan(&uid, &username, &password)
	if err != nil {
		fmt.Println("查询出现异常=", err)
		return nil, err
	}
	u := &User{
		Uid:      uid,
		Username: username,
		Password: password,
	}
	return u, nil
}
func (user *User) GetUserByUid() (*User, error) {
	sqlStr := "select uid, username, password from user where uid = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, user.Uid)
	//声明变量
	var (
		uid      int
		username string
		password string
	)
	err := row.Scan(&uid, &username, &password)
	if err != nil {
		fmt.Println("查询出现异常=", err)
		return nil, err
	}
	u := &User{
		Uid:      uid,
		Username: username,
		Password: password,
	}
	return u, nil
}

//获取所有的数据库用户
func (user *User) GetUsers() ([]*User, error) {
	sqlStr := "select uid, username, password from user"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var users []*User
	for rows.Next() {
		//声明变量
		var (
			uid      int
			username string
			password string
		)
		err := rows.Scan(&uid, &username, &password)
		if err != nil {
			fmt.Println("查询出现异常=", err)
			return nil, err
		}
		u := &User{
			Uid:      uid,
			Username: username,
			Password: password,
		}
		users = append(users, u)
	}
	return users, nil
}
