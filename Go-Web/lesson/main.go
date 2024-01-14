package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//定义模版
	//解析模版
	t, err := template.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("解析模版错误：%v", err)
		return
	}
	//渲染模版
	u1 := User{
		Name:   "小王子",
		Gender: "男",
		Age:    22,
	}
	m1 := map[string]interface{}{
		"name":   "xwz",
		"gender": "nan",
		"age":    122,
	}
	t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP服务错误：%v", err)
		return
	}
}
