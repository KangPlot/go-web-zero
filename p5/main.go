package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		// 打印header
		fmt.Fprintln(w, r.Header)

		// 打印body
		length := r.ContentLength
		body := make([]byte, length)
		r.Body.Read(body)
		fmt.Fprintln(w, string(body)) //  结果写入到ResponseWrite 里面，需要转为字符串，才能看明白

	})
	server.ListenAndServe()
}
