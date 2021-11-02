package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"ssevven/model"
	"strconv"
	"strings"
)

//show /login.html
func Login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./client/templates/login.html")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}
	args := make(map[string]interface{})
	args["Home"] = model.Cfg.Domain
	args["RegistAction"] = model.Cfg.Domain + "/regist"
	args["LoginAction"] = model.Cfg.Domain + "/loginProcess"
	t.Execute(w, args)
}

//process login request
func LoginProcess(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println()
	log.Println("login:", r.PostForm)
	user := &model.User{}
	user.Username = r.PostForm["username"][0]
	user.Password = r.PostForm["password"][0]
	u, _ := user.GetUserByUsername()
	// log.Println("u=", u)
	res := &model.LoginRes{}
	if u == nil {
		res.Code = 310 //用户名不存在
		res.Msg = "用户名不存在"
	} else {
		if u.Password == user.Password {
			res.Code = 300
			res.Msg = "登陆成功"
			res.UserUrl = model.Cfg.Domain + "/user/" + strconv.Itoa(u.Uid)
			log.Printf("登录成功=%#v\n", user)
		} else {
			res.Code = 320 //密码错误
			res.Msg = "密码错误"
		}
	}
	// Set response header
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)
	log.Println("res =", res)
}

//show /regist.html
func Regist(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./client/templates/regist.html")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}
	args := make(map[string]interface{})
	args["Login"] = model.Cfg.Domain + "/login"
	args["RegistAction"] = model.Cfg.Domain + "/registProcess"
	t.Execute(w, args)
}

//process regist request
func RegistProcess(w http.ResponseWriter, r *http.Request) {
	//如果username在数据库中差找不到，则允许注册
	//否则，返回注册失败
	r.ParseForm()
	fmt.Println()
	log.Println("regist:", r.PostForm)
	user := &model.User{}
	user.Username = r.PostForm["username"][0]
	user.Password = r.PostForm["password"][0]
	u, _ := user.GetUserByUsername()
	// log.Println("u=", u)
	res := &model.RegistRes{}

	if u != nil {
		res.Code = 220 //用户已注册
		res.Msg = "用户名已注册"
	} else {
		err := user.AddUser()
		if err != nil {
			log.Println("dml err=", err)
			return
		}
		res.Code = 200 //注册成功
		res.Msg = "注册成功"
		log.Printf("注册成功=%#v\n", user)
	}
	// Set response header
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)
	log.Println("res =", res)
}

//直接访问//Todo
//show /user.html
func UserProcess(w http.ResponseWriter, r *http.Request) {
	args := make(map[string]interface{})
	user := &model.User{}
	user.Uid, _ = strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
	u, _ := user.GetUserByUid()
	if u == nil {
		t, err := template.ParseFiles("./client/templates/index.html")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err)
			return
		}
		t.Execute(w, nil)
	} else {
		t, err := template.ParseFiles("./client/templates/user.html")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err)
			return
		}
		args["GameOneRequest"] = model.GameOne.GameURL
		args["GameOneName"] = model.GameOne.GameName
		args["GameTwoRequest"] = model.GameTwo.GameURL
		args["GameTwoName"] = model.GameTwo.GameName
		args["UserImg"] = u.Username
		args["UserName"] = u.Username
		args["Login"] = model.Cfg.Domain + "/login"
		t.Execute(w, args)
	}
}
