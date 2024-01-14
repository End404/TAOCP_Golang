package main

import (
	"fmt"
	"strings"
)

/*第3章 复合数据类型
3.2 map
*/

func main() {
	// var a map[string]int //声明map类型，没有初始化
	// fmt.Println(a == nil)
	// a = make(map[string]int, 8) //map的初始化
	// fmt.Println(a == nil)
	// // map中添加键值对
	// a["shnz"] = 100
	// a["shxwz"] = 200
	// fmt.Println(a)
	// fmt.Printf("a: %#v\n", a)
	// fmt.Printf("type:%T\n", a)
	// //声明的同时完成初始化
	// b := map[int]bool{
	// 	1: true,
	// 	2: false,
	// }
	// fmt.Printf("b:%#v\n", b)
	// fmt.Printf("type:%T", b)

	//判断某个键存不存在
	//遍历map
	// var scoreMap = make(map[string]int, 8)
	// scoreMap["shnz"] = 100
	// scoreMap["shxwz"] = 200

	//判断 zegz 在不在 scoreMap中
	// value, ok := scoreMap["zegz"]
	// if ok {
	// 	fmt.Println("在scoreMap中", value)
	// } else {
	// 	fmt.Println("查无此人")
	// }

	//遍历
	// for range
	// map是无序的
	// for k, v := range scoreMap {
	// 	fmt.Println(k, v)
	// }
	//遍历key
	// for k := range scoreMap {
	// 	fmt.Println(k)
	// }
	//遍历value
	// for _, v := range scoreMap {
	// 	fmt.Println(v)
	// }

	//删除键值对
	// delete(scoreMap, "shxwz")
	// fmt.Println(scoreMap)

	//按照某个固定循序遍历map
	// var scoreMap_1 = make(map[string]int, 100)
	// for i := 0; i < 50; i++ { //添加50个键值对
	// 	key := fmt.Sprintf("stu%02d", i)
	// 	value := rand.Intn(100) //0~99随机整数
	// 	scoreMap_1[key] = value
	// }
	// for k, v := range scoreMap_1 {
	// 	fmt.Println(k, v)
	// }

	//按照key从小到大的循序遍历scoreMap_1
	// 1.先取出所有的key存放到切片中
	// keys := make([]string, 0, 100)
	// for k := range scoreMap_1 {
	// 	keys = append(keys, k)
	// }
	// // 2. 对key做排序
	// sort.Strings(keys)
	// // 3. 按照排序后的key对scoreMap_1排序
	// for _, key := range keys {
	// 	fmt.Println(key, scoreMap_1[key])
	// }

	// 元素类型为map的切片
	// var mapSlice = make([]map[string]int, 8, 8) //支持完成了切片的初始化
	// //  需要完成内部map元素的初始化
	// mapSlice[0] = make(map[string]int, 8) //完成了map的初始化
	// mapSlice[0]["扇席温枕"] = 100
	// fmt.Println(mapSlice)

	//值为切片
	// var sliceMap = make(map[string][]int, 8)
	// v, ok := sliceMap["中国"]
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	sliceMap["中国"] = make([]int, 8)
	// 	sliceMap["中国"][0] = 100
	// 	sliceMap["中国"][1] = 200
	// 	sliceMap["中国"][2] = 300
	// }
	// for k, v := range sliceMap { //遍历sliceMap
	// 	fmt.Println(k, v)
	// }

	//统计单词出现的次数
	// "how do you do"单词出现次数
	var s = "how do you do" //定义一个map
	var wordCount = make(map[string]int, 10)
	words := strings.Split(s, " ")
	for _, word := range words { //遍历单词做统计
		v, ok := wordCount[word]
		if ok {
			// map中有这个单词的统计记录
			wordCount[word] = v + 1
		} else {
			//map中没有这个单词的统计记录
			wordCount[word] = 1
		}
	}
	for k, v := range wordCount {
		fmt.Println(k, v)
	}

}
