package controller

import (
	"fmt"
	"net/http"
	"text/template"
)

// 注册路由
func registerContactRoutes() {
	http.HandleFunc("/contact", handleContact)
}

// handle处理逻辑
func handleContact(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("p21/layout.html", "p21/home.html")
	if err != nil {
		fmt.Printf("错误%v", err)                       // Ugly debug output
		w.WriteHeader(http.StatusInternalServerError) // Proper HTTP response
		return
	}
	t.ExecuteTemplate(w, "layout", "hello world")
}
