package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

func Game1(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./client/templates/huarongdao.html")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}
	t.Execute(w, nil)
}
func Game2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./client/templates/huarongdao.html")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}
	t.Execute(w, nil)
}
