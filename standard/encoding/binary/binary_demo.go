package main

import "encoding/binary"

type Head struct {
	Cmd     byte
	Version byte
	Magic   uint16
	Reserve byte
	HeadLen byte
	BodyLen uint16
}

//这个是一个常见的在tcp 拼包得例子。在例子中通过binary.BigEndian.Uint16将数据按照网络序的格式读出来，放入到head中对应的结构里面。
func NewHead(buf []byte) *Head {
	head := new(Head)

	head.Cmd = buf[0]
	head.Version = buf[1]
	head.Magic = binary.BigEndian.Uint16(buf[2:4])
	head.Reserve = buf[4]
	head.HeadLen = buf[5]
	head.BodyLen = binary.BigEndian.Uint16(buf[6:8])
	return head
}
