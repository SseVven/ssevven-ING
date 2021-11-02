package main

import (
	"log"
	"net/http"
	"ssevven/controller"
	"ssevven/model"
)

func init() {
	//读取服务器配置文件
	var err error
	model.Cfg, err = model.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}
	//预编译所有的模板
	// model.LoadTemplates()
	//注册处理器
	controller.RegisterHandlers()
	//测试连接数据库
	user := &model.User{}
	users, err := user.GetUsers()
	if err != nil {
		log.Println("GetUsers() err =", err)
	}
	for i, v := range users {
		log.Printf("user %v : %#v\n", i, v)
	}
}
func main() {
	//设置web服务器
	server := http.Server{
		Addr:    model.Cfg.AppHost + model.Cfg.AppPort,
		Handler: nil, //DefaultServeMux
	}
	server.ListenAndServe()
	// server.ListenAndServeTLS("cert.pem", "key.pem")
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fileName := r.URL.Path[1:]
	// 	t := template.Lookup(fileName)
	// 	if t != nil {
	// 		var err error
	// 		if fileName == "user.html" {
	// 			err = t.Execute(w, map[string]string{
	// 				"jack": "男;12岁;在线;你在干嘛",
	// 				"mary": "女;17岁;离线;さようなら！",
	// 			})
	// 		} else {
	// 			err = t.Execute(w, nil)
	// 		}
	// 		if err != nil {
	// 			log.Fatalln(err.Error())
	// 		}
	// 	} else {
	// 		w.WriteHeader(http.StatusNotFound)
	// 	}
	// })
	// http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
	// 	t, _ := template.ParseFiles("templates/layout.html", "templates/login.html")
	// 	t.ExecuteTemplate(w, "layout", "login")
	// })
	// http.Handle("/css/", http.FileServer(http.Dir("assets")))
}
