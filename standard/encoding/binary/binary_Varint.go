package main

import (
	"encoding/binary"
	"fmt"
)

//Varint 从 buf 解码 int64 并返回该值和读取的字节数（> 0）。如果发生错误，则值为0，并且字节数 n <= 0，其含义如下：
//n == 0：buf太小了
//n <0：大于64位的值（溢出）
//和-n是读取的字节数

func main() {
	inputs := [][]byte{
		[]byte{0x81, 0x01},
		[]byte{0x7f},
		[]byte{0x03},
		[]byte{0x01},
		[]byte{0x00},
		[]byte{0x02},
		[]byte{0x04},
		[]byte{0x7e},
		[]byte{0x80, 0x01},
	}
	for _, b := range inputs {
		x, n := binary.Varint(b)
		if n != len(b) {
			fmt.Println("Varint did not consume all of in")
		}
		fmt.Println(x)
	}
}
