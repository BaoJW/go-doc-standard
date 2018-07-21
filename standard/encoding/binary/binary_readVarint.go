package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//从r读取一个编码后的有符号整数，并返回该整数
func main() {
	var emptyBuf []byte
	buf := []byte{144, 192, 192, 129, 132, 136, 140, 144, 16, 0, 1, 1}
	overFlowBuf := []byte{144, 192, 192, 129, 132, 136, 140, 144, 192, 192, 1, 1}

	num, err := binary.ReadVarint(bytes.NewBuffer(emptyBuf))
	fmt.Println(num, err)

	num, err = binary.ReadVarint(bytes.NewBuffer(buf))
	fmt.Println(num, err)

	num, err = binary.ReadVarint(bytes.NewBuffer(overFlowBuf))
	fmt.Println(num, err)
}

//类型名称        字节数    取值范围
//signed char       1        -128～+127
//short int         2        -32768～+32767
//int               4        -2147483648～+2147483647
//long int          4        -2147483648～+2141483647
//long long long int  8      -9223372036854775808～+9223372036854775807

//unsigned int  （unsigned long）
//4字节8位可表达位数：2^32=42 9496 7296
//范围：0 ～ 42 9496 7295 (42*10^8)
//
//int  （long）
//4字节8位可表达位数：2^32=42 9496 7296
//范围：-21 4748 3648 ～ 21 4748 3647 (21*10^8)
//
//long long (__int64)
//8字节8位可表达位数：2^64=1844 6744 0737 0960 0000
//范围：-922 3372 0368 5477 5808 ～ 922 3372 0368 5477 5807 (922*10^16)
//
//unsigned long (unsigned __int64)
//8字节8位可表达位数：2^64=1844 6744 0737 0960 0000
//范围：0 ～ 1844 6744 0737 0955 1615 (1844*10^16)
