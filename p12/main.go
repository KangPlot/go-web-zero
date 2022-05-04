package main

import (
	"html/template"
	"net/http"
)

// 自定义的handler
func process(w http.ResponseWriter, r *http.Request) {
	// 解析模板文件,相对路径
	t, _ := template.ParseFiles("tmpl.html")
	// 执行模板，传入数据，替换{{.}}
	t.Execute(w, "hello go go go go")
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/test", process)
	server.ListenAndServe()
}
