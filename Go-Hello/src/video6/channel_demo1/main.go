package main

import "fmt"

/*
	两个goroution，两个channel
	1.生产0~100的数字发送到ch1。
	2.从ch1中取出数据计算他的平方，把结果发送到ch2中。
*/

func f1(ch chan<- int) { //生产0到100数字
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}
func f2(ch1 <-chan int, ch2 chan<- int) { //取出数字计算平方，发送到ch2中
	for false { //从通道中取值的方式1
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}
func main() {

	ch1 := make(chan int, 100)
	ch2 := make(chan int, 200)

	go f1(ch1)
	go f2(ch1, ch2)

	for ret := range ch2 { //通道取值方式2
		fmt.Println(ret)
	}
}
