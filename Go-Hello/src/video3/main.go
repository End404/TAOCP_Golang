package main

import (
	"fmt"
	"time"
)

//time demo

func main() {
	now := time.Now() //时间对象
	/*
		fmt.Println(now)
		year := now.Year()
		month := now.Month()
		day := now.Day()
		fmt.Println(year, month, day)
		//时间戳
		timeStamp1 := now.Unix()
		timeStamp2 := now.UnixNano()
		fmt.Println(timeStamp1, timeStamp2)
		//时间戳转换具体的时间格式
		t := time.Unix(156150154, 0)
		fmt.Println(t)
		//时间间隔
		n := 5
		time.Sleep(time.Duration(n) * time.Second)
		fmt.Println("over")
	*/
	t2 := now.Add(time.Hour)
	fmt.Println(t2)
	fmt.Println(t2.Sub(now))

	//定时器
	//for tmp := range time.Tick(time.Second * 2) {
	//	fmt.Println(tmp)
	//}

	//时间格式化
	ret1 := now.Format("2006-01-05")
	fmt.Println(ret1)

	//解析字符串类型的时间
	timeStr := "2019/08/07 15:00:00"
	//1.拿到时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	//2.根据时区区解析一个字符串格式的时间
	timeObj, err := time.Parse("2006/01/02 15:04:05", timeStr)
	if err != nil {
		fmt.Println("parse timeStr failed,err: %v\n", err)
		return
	}
	fmt.Println(timeObj)
	timeObj2, err := time.ParseInLocation("2006/01/02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Println("parse timeStr failed,err: %v\n", err)
		return
	}
	fmt.Println(timeObj2)

}
