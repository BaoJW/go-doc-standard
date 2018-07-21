package main

import (
	"encoding/binary"
	"fmt"
)

//将一个int64数字编码写入buf并返回写入的长度，如果buf太小，则会panic。
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("server is panic.")
		}
	}()
	buf := make([]byte, binary.MaxVarintLen64)

	//test panic
	//buf := make([]byte, 1)

	for _, x := range []int64{-65, -64, -2, -1, 0, 1, 2, 63, 64} {
		n := binary.PutVarint(buf, x)
		fmt.Printf("%x\n", buf[:n]) //%x十六进制表示，字母形式为小写a-f
	}
}
