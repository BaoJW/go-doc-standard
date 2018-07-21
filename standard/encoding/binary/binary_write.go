package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

//写将数据的二进制表示写入 w。数据必须是固定大小的值或固定大小的值片段，或指向这些数据的指针。布尔值编码为一个字节：1 为真，0 为假。
// 写入 w 的字节使用指定的字节顺序进行编码，并从数据的连续字段中读取。在编写结构时，将为具有空白（_）字段名称的字段写入零值。

func main() {
	buf := new(bytes.Buffer)
	var pi float64 = math.Pi
	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", buf.Bytes())
}
