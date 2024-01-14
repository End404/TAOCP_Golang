package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模版
	t, err := template.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("解析模版错误：%v", err)
		return
	}
	//渲染模版
	name := "xiaowangzi"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("渲染模版错误：%v", err)
		return
	}

}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP服务错误：%v", err)
		return
	}
}
