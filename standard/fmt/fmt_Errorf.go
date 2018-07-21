package main

import "fmt"

//test Errorf
//Errorf格式根据格式说明符并返回字符串作为满足错误的值。
func main()  {
	//相当于格式化错误
	err := fmt.Errorf("invalid error %v", "hahhaha")
	fmt.Println(err)
}
