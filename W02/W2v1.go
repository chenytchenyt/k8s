package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 记录访问日志
	log.Printf("Client IP: %s, Response Code: %d\n", r.RemoteAddr, http.StatusOK)
	// 从环境变量中读取 VERSION 配置
	version := os.Getenv("VERSION")
	// 如果 VERSION 配置存在，则将其添加到响应头中
	if version != "" {
		w.Header().Set("Version", version)
	}
	// 遍历请求头并把它们添加到响应头
	for name, values := range r.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}
	// 向响应写入状态码和消息
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, World!")
}
