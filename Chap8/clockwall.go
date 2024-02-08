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
	go DialClockServe("8000", &timewall[0])
	go DialClockServe("8001", &timewall[1])
	go DialClockServe("8002", &timewall[2])
	printwall(&timewall)
}
func DialClockServe(port string, dst *string) {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(dst, conn)
}
func mustCopy(dst *string, src io.Reader) {
	for {
		if _, err := fmt.Fscanln(src, dst); err != nil {
			log.Fatal(err)
		}
	}
}
func printwall(timewall *[3]string) {
	for {
		fmt.Println(timewall[0] + " || " + timewall[1] + " | " + timewall[2])
		time.Sleep(1 * time.Second)
	}
}
