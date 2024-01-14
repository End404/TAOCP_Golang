package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//解析模版
	t, err := template.New("index.tmpl").
		Delims("{[", "]}").
		ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("解析错误：", err)
		return
	}
	//渲染模版
	name := "xwz"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("渲染错误：", err)
		return
	}
}
func xss(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./xss.tmpl")
	str := "<script></script>"
	t.Execute(w, str)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP服务错误：%v", err)
		return
	}
}
