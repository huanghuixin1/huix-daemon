package main

import (
	"os"
	"fmt"
)

func main() {
	file, err := os.Open("./config.conf")
	defer file.Close()
	fileInfo, _ := file.Stat()
	if (err != nil) {
		panic(err)
	}
	fileSize := fileInfo.Size()
	b := make([]byte,fileSize)
	len, errFile := file.Read(b)
	if (errFile != nil) {
		panic(errFile)
	}
	file.Close()
	fmt.Println(string(b), "输出结果", len)
}
