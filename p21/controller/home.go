package controller

import (
	"log"
	"net/http"
	"text/template"
)

// 注册路由
func registerHomeRoutes() {
	http.HandleFunc("/home", handleHome)
}

// handle处理逻辑
func handleHome(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("p21/layout.html", "p21/home.html")
	log.Println(e)
	t.ExecuteTemplate(w, "layout", "hello world")
}
