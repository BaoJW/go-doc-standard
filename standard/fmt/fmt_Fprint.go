package main

import (
	"os"
	"bufio"
	"fmt"
)

//test Fprint
//格式化写入文件你想写的内容
func main()  {
	// Create a file and use bufio.NewWriter.
	//这里每次执行程序时都会创建一次Fprint.txt，所以你会发现每次写入文本的内容都是会覆盖掉上次的内容
	f, _ := os.Create("./standard/fmt/Fprint.txt")
	w := bufio.NewWriter(f)

	// Use Fprint to write things to the file.
	// ... No trailing newline is inserted.
	fmt.Fprint(w, "Hello")
	fmt.Fprint(w, " World")
	fmt.Fprint(w, " 123")
	fmt.Fprint(w, "...")
	fmt.Fprint(w, "\n")

	// Use Fprintf to write formatted data to the file.
	EnglishName := "lavine"
	ChineseName := "baojiawei"
	fmt.Fprintf(w, "%v %v\n", "My english name is", EnglishName)
	fmt.Fprintf(w, "%v %v\n", "My chinese name is", ChineseName)


	w.Flush()
}