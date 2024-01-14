package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int, results chan int) {	//池并发
	for p := range ports{
		address := fmt.Sprintf("20.194.168.28:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {	

	ports := make(chan int, 100) 
	results := make(chan int)
	var openports []int
	var closedports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i < 1024; i++ {
			ports <- i
		}
	}()

	for i := 1; i < 1024; i++ {
		port := <- results
		if port != 0 {
			openports = append(openports, port)
		} else {
			closedports = append(closedports, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openports)
	sort.Ints(closedports)

	for _, port := range closedports {
		fmt.Printf("%d closed\n", port)
	}
	
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}

	/*
func main(){
	//并发
	start := time.Now()
	var wg sync.WaitGroup
	for i := 21; i < 120; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("20.194.168.28:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%s 关闭了\n", address)
				return
			}
			conn.Close()
			fmt.Printf("%s 打开了！！！\n", address)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("\n\n%d seconds", elapsed)
}
*/

// func main() { //非并发
// 	for i := 21; i < 120; i++ {
// 		address := fmt.Sprintf("20.194.168.28:%d", i)
// 		conn, err := net.Dial("tcp", address)
// 		if err != nil {
// 			fmt.Printf("%s 关闭了\n", address)
// 			continue
// 		}
// 		conn.Close()
// 		fmt.Printf("%s 打开了！！！\n", address)
// 	}
// }

