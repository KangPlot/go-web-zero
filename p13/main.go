package main

import (
	"net/http"
	"text/template"
)

func main() {

	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()

}

func process(w http.ResponseWriter, r *http.Request) {
	//解析1个模板
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, "hello go")
	////等价于下面两句话
	//t := template.New("tmpl.html")
	//t, _ = t.ParseFiles("tmpl.html")

	//t1，_, := template.ParseGlob("*.html")

	//解析多个模板
	ts, _ := template.ParseFiles("t1.html", "t2.html")
	ts.ExecuteTemplate(w, "t2.html", "hello go t2")

}
