package controller

import (
	"net/http"
)

func RegisterHandlers() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/loginProcess", LoginProcess)
	http.HandleFunc("/regist", Regist)
	http.HandleFunc("/registProcess", RegistProcess)
	//Todo
	http.HandleFunc("/user/", UserProcess)
	//todo
	http.HandleFunc("/game/game1", Game1)
	http.HandleFunc("/game/game2", Game2)
	//tODO
	http.HandleFunc("/chatroom/", ChatroomProcess)
	// http.Handle("../", http.StripPrefix("../", http.FileServer(http.Dir("./client"))))
	// http.Handle("../assets/", http.StripPrefix("../assets/", http.FileServer(http.Dir("./client"))))
}
