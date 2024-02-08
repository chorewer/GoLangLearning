// 由于只想要单文件程序 这里会报错 不用管
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	var timewall = [3]string{"America Time ", "Shanghai Time", "Paris Time"}
	go DialClockServe("8000", &timewall[0]) //端口已经写死了，也可以根据实际情况写入任意port
	go DialClockServe("8001", &timewall[1])
	go DialClockServe("8002", &timewall[2])
	printwall(&timewall)
}
func DialClockServe(port string, dst *string) {
	//使用多个协程向多个port请求数据
	//Dial库函数会提供一个conn，其实现了io.Reader接口
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(dst, conn)
}
func mustCopy(dst *string, src io.Reader) {
	for {
		//针对每一个异步地更新每一个string
		if _, err := fmt.Fscanln(src, dst); err != nil {
			log.Fatal(err)
		}
	}
}
func printwall(timewall *[3]string) {
	for {
		//每秒钟根据当前string更新时钟
		fmt.Printf("\r" + timewall[0] + " || " + timewall[1] + " || " + timewall[2])
		time.Sleep(1 * time.Second)
	}
}
