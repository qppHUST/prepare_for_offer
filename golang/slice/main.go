package main

import "fmt"

func main() {
	// slice是不定长的 ,array是定长的
	// slice和array都是按值传递的  ,地址也是值
	// var slice []int
	// var array [10]int //长度是编译时定义
	//slice的扩容： 小于1024的时候乘2，大于的时候乘1.25
	//index赋值优于append
	//make的时候尽量预先分配内存
	//如果能确定访问到slice边界，
	//那么最好可以开始的时候就访问一下边界，这样之后访问其他的下标的时候就不会去检查是否下标越界了，称为bce

	// var s []int
	// // ss := []int{}   //这样初始化更好
	// for i := 3; i < 1025; i++ {
	// 	s = append(s, i)
	// }
	// fmt.Println(len(s)) //1025
	// fmt.Println(cap(s)) //1280

	//slice删除元素
	sslice := []int{0, 1, 2, 3, 4}
	sslice = append(sslice[:0], sslice[1:]...)
	fmt.Printf("%v", sslice)
}

// runtime.go
// type slice struct {
// 	array unsafe.Pointer
// 	len   int
// 	cap   int
// }
