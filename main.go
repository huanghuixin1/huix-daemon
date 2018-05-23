package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func main() {
	for true  {
		fmt.Println("无线打印")
		time.Sleep(time.Second * 5)
	}

	fmt.Println("开始执行")
	ret := exec.Command("bash", "-c", "ps -ef | grep frps_ssh")
	//ret := exec.Command("ipconfig")
	retBytes, err := ret.Output()
	if (err != nil) {
		fmt.Println(err, "报错！！！")
		return
	}
	fmt.Println("执行后的输出如下")
	retStr := string(retBytes)
	fmt.Println(retStr)

	if (strings.Index(retStr, "/usr/local/frp/frps -c /usr/local/frp/frps_ssh.ini") > -1 || strings.Index(retStr, "frps -c frps_ssh.ini") > -1) {
		fmt.Println("存在ssh frp进程")
	} else {
		fmt.Println("不存在frp进程 需要重新启动")
		retFrp := exec.Command("bash", "-c", "nohup /usr/local/frp/frps -c /usr/local/frp/frps_ssh.ini > myout.file 2>&1 &")
		retFrpBytes, errFrp := retFrp.Output()

		if (errFrp != nil) {
			fmt.Println("出现错误", string(retFrpBytes), errFrp.Error())
		}
	}

	fmt.Println("执行完毕")
}
