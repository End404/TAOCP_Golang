package main

import (
	"fmt"
	"sync"
)

//并发编程

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello 娜扎", i)
	wg.Done() //通知计数器-1
}

func main() {

	wg.Add(10) //计数牌+1
	for i := 0; i < 10; i++ {
		go hello(i) //开启并发线程
	}
	fmt.Println("hello 主函数")
	//time.Sleep(time.Second)
	wg.Wait() //阻塞；等待
}
