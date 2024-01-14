package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var ( //自定义log
	Trace   *log.Logger //几乎任何东西
	Info    *log.Logger //重要信息
	Warning *log.Logger //警告
	Error   *log.Logger //错误
)

func init() {

	/*
		log.SetPrefix("CC: ")

		f, err := os.OpenFile("./yx.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY,
			0666)
		if err != nil {
			log.Fatalln(err)
		}
		log.SetOutput(f)

		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	*/

	//自定义logger
	file, err := os.OpenFile("errors.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("无法打开错误 log 文件：", err)
	}

	Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {
	/*
		// log.Println("1234")
		// log.Fatalln("1234")
		// log.Panicln("1234")
	*/

	Trace.Println("小事")
	Info.Println("信息")
	Warning.Println("警告")
	Error.Println("故障")
}
