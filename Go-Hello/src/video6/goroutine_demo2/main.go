package main

import (
	"fmt"
	"sync"
)

//并发编程

var wg sync.WaitGroup

func main() {

	wg.Add(10) //计数牌+1
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("hello", i)
			wg.Done()
		}(i) //开启并发线程
	}
	fmt.Println("hello 主函数")
	wg.Wait() //阻塞；等待
}
