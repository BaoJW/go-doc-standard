package main

import (
	"encoding/binary"
	"fmt"
)

//Uvarint 从 buf 解码 uint64 并返回该值和读取的字节数（> 0）。如果发生错误，则该值为0，并且字节数n <= 0意味着：
//n == 0：buf太小了
//n <0：大于64位的值（溢出）
//和-n是读取的字节数
func main() {
	inputs := [][]byte{
		[]byte{0x01},
		[]byte{0x02},
		[]byte{0x7f},
		[]byte{0x80, 0x01},
		[]byte{0xff, 0x01},
		[]byte{0x80, 0x02},
	}
	for _, b := range inputs {
		x, n := binary.Uvarint(b)
		if n != len(b) {
			fmt.Println("Uvarint did not consume all of in")
		}
		fmt.Println(x)
	}
}
