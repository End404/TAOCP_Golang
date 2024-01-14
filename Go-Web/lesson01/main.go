package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//定义函数
	k := func(name string) (string, error) {
		return name + "年轻", nil
	}

	//解析模版
	t := template.New("f.tmpl") //创建模版对象
	t.Funcs(template.FuncMap{
		"kua": k,
	})
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Println("解析模版错误：", err)
		return
	}
	name := "xwz"
	t.Execute(w, name) //渲染模版
}

func main() {
	http.HandleFunc("/", f1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP服务错误：%v", err)
		return
	}
}
