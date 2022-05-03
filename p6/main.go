package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL
		query := url.Query()

		// 方法1：通过map，返回字符串切片
		id := query["id"]
		log.Println(id)

		// 方法2：通过Get方法，返回key对应的第一个值
		name := query.Get("name")
		log.Println(name)

	})
	http.ListenAndServe("localhost:8080", nil)
}

// 在浏览器输入：http://localhost:8080/home?id=001&name=xpk&id=002&name=zhuanma
