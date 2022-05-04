package controller

import (
	"fmt"
	"net/http"
	"text/template"
)

// 注册路由
func registerAboutRoutes() {
	http.HandleFunc("/about", handleAbout)
}

// handle处理逻辑
func handleAbout(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("p21/layout.html", "p21/about.html")
	// 服务器没有数据，需要判err, open layout.html: The system cannot find the file specified.
	if err != nil {
		fmt.Println(err)                              // Ugly debug output
		w.WriteHeader(http.StatusInternalServerError) // Proper HTTP response , HTTP ERROR 500
		return
	}
	t.ExecuteTemplate(w, "about", "hello world")
}
