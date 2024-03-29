package main

import (
	"fmt"
	"time"
)

//work_pool

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println(id, job)
		results <- job * 2
		time.Sleep(time.Microsecond * 500)
		fmt.Println(id, job)
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//开启3个并发
	for j := 0; j < 3; j++ {
		go worker(j, jobs, results)
	}

	//发送5个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)
	//输出结果
	for i := 0; i < 5; i++ {
		ret := <-results
		fmt.Println(ret)
	}
}
