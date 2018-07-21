package main

import (
	"encoding/binary"
	"fmt"
)

//Size 返回 Write 将生成的值，以便对值 v 进行编码，该值必须是固定大小的值或固定大小的值片段或指向此类数据的指针。如果 v 不是这些，则 Size 返回-1。
func main() {
	var (
		intNum     int8 //固定大小的值指的是像这样的，不是int这样的，不确定是哪种int类型
		int32Num   int32
		int64Num   int64
		float32Num float32
		float64Num float64
		stringNum  string //数字通过读取和写入固定大小的值进行翻译。固定大小的值可以是固定大小的算术类型（bool，int8，uint8，int16，float32，complex64，...），也可以是只包含固定大小值的数组或结构。
	)
	intNum = 100
	int32Num = 101
	int64Num = 102
	float32Num = 3.1415925
	float64Num = 0.1415926
	stringNum = "hello"

	intSize := binary.Size(intNum)
	int32Size := binary.Size(int32Num)
	int64Size := binary.Size(int64Num)
	float32Size := binary.Size(float32Num)
	float64Size := binary.Size(float64Num)
	stringSize := binary.Size(stringNum)

	fmt.Println("intSize is", intSize)
	fmt.Println("int32Size is", int32Size)
	fmt.Println("int64Size is", int64Size)
	fmt.Println("float32Size is", float32Size)
	fmt.Println("float64Size is", float64Size)
	fmt.Println("stringSize is", stringSize)

}
