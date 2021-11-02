package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"ssevven/model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./client/templates/index.html")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}
	args := make(map[string]interface{})
	args["LoginHref"] = model.Cfg.Domain + "/login"
	t.Execute(w, args)
}
