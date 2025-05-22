package main

import "fmt"

func updateAge(age *int) {
	*age = 30
}

func main() {
	a := 10
	updateAge(&a)  // &a 取地址传入
	fmt.Println(a) // 输出 30
}
