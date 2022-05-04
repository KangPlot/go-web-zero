package main

import (
	"fmt"
	"go-web-zero/p21/controller"
	"net/http"
)

func main() {

	/* 写在controller 里面
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("layout.html", "home.html")
		t.ExecuteTemplate(w, "layout", "hello world")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("layout.html", "about.html")
		t.ExecuteTemplate(w, "layout", "hello world")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		t, e := template.ParseFiles("layout.html")
		t.ExecuteTemplate(w, "layout", "hello world")
		log.Println(e)
	})
	*/
	fmt.Println("准备注册路由")
	controller.RegisterRoutes() // 注册路由，等价于上面的代码
	fmt.Println("已经注册路由")
	http.ListenAndServe("localhost:8080", nil) // 采用默认的多路复用器
}

// 问题，ｗｅｂ　无任何内容？？？
