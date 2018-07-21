package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//binary:二进制的意思
//从r中读取binary编码的数据并赋给data,data必须是一个指向定长值的指针或者定长值的切片。
//从r读取的字节使用order指定的字节序编码并写入data的字段里。名字中有'_'的字段会被跳过，这些字段可用于填充（内存空间）。
func main() {
	exampleOne()
	exampleTwo()
}

func exampleOne() {
	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Print(pi)
}

func exampleTwo() {
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40, 0xff, 0x01, 0x02, 0x03, 0xbe, 0xef}
	r := bytes.NewReader(b)

	var data struct {
		PI   float64
		Uate uint8
		Mine [3]byte
		Too  uint16
	}

	if err := binary.Read(r, binary.LittleEndian, &data); err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	fmt.Println(data.PI)
	fmt.Println(data.Uate)
	fmt.Printf("% x\n", data.Mine)
	fmt.Println(data.Too)
}
