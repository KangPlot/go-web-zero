package middleware

import "net/http"

// 链式结构， Next 设置为 什么，下一个handler 就是什么
// AuthMiddleware ..
type AuthMiddleware struct {
	Next http.Handler
}

func (am *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 如果只有一个中间件，改中间件的字段next 为nil, 交给默认路由器处理
	if am.Next == nil {

		am.Next = http.DefaultServeMux
	}
	// 判断auth
	auth := r.Header.Get("Authorization")
	if auth != "" {
		// before 路由
		am.Next.ServeHTTP(w, r)
		// after 路由
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
