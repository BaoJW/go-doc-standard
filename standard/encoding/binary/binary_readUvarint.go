package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//从r读取一个编码后的无符号整数，并返回该整数
//如果读取的值为有符号的则会报EOF
func main() {
	i := int64(-100)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, i)
	fmt.Println("after write, buf is:", buf.Bytes())

	num, err := binary.ReadUvarint(buf)
	if err != nil {
		fmt.Println("error is", err)
		return
	}

	fmt.Println("read uvarint num is:", num)
}
