package main

import "net/http"

func main() {
	mh := myHandler{} // 要使用的是这个指针
	about := aboutHandler{}
	hello := helloHandler{}

	//http.ListenAndServe("localhost:8080", nil)

	// 等价于上面一句话
	server := http.Server{
		Addr:    ("localhost:8080"),
		Handler: nil,
		//Handler: &mh,
	}
	//不同路径对应不用handler
	http.Handle("/hello", &hello)
	http.Handle("/about", &about)
	http.Handle("/home", &mh)

	server.ListenAndServe()

}

// 自定义handler, 实现ServerHTTP 方法
type myHandler struct {
}

// 自定义handler, 实现ServerHTTP 方法
type helloHandler struct {
}

type aboutHandler struct {
}

//ServeHTTP 不是ServerHTTP
func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home handler"))
}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello handler"))
}
func (m *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about handler"))
}
