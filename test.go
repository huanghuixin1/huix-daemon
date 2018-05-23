package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	fmt.Println(strings.HasSuffix("111.conf", ".conf"))
	fmt.Println(strings.HasSuffix("111.conf.demo", ".conf"))
	fmt.Println(strings.HasSuffix("111.conf.11", ".conf"))
	files, _ := ioutil.ReadDir("./process")
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
