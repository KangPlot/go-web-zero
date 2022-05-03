package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		// r.Form 出现情况，url 里面的key 和 表单里面相同的话，会一起显示出来
		fmt.Fprintln(w, r.PostForm)
	})
	server.ListenAndServe()
}
