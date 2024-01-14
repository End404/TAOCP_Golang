package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func inddex(w http.ResponseWriter, r *http.Request) {
	//解析模版
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("解析模版错误:", err)
		return
	}
	msg := "xwz"
	t.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	//解析模版
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Println("解析模版错误:", err)
		return
	}
	msg := "小王子"
	t.Execute(w, msg)
}

func main() {
	http.HandleFunc("/index", inddex)
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP服务错误：%v", err)
		return
	}
}
