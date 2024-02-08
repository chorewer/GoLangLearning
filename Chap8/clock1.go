// 由于只想要单文件程序 这里会报错 不用管
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func main() {
	//使用flag包去获取到命令行的标签作为端口号
	// -port=8000
	var ip = flag.Int("port", 8001, "give the prog a port")
	//用于模拟该clock给出的时间的时区，模拟有许多分布的时间服务器，运行再不同的端口上
	// -postion="America/New_York|Asia/Shanghai|Europe/Paris" 具体配置在 $GOROOT/lib/time/zoneinfo.zip
	var address = flag.String("postion", "Shanghai", "give the prog a address")

	flag.Parse() //定义结束之后去解析命令行标签
	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(*ip))
	fmt.Println("localhost:" + strconv.Itoa(*ip))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		postion, locerr := time.LoadLocation(*address)

		if locerr != nil {
			log.Fatal(err)
		}
		go handleConn(conn, postion)
		//go handleConn(conn)
		//如果 不使用 goroutine的话 每一个连接都会阻塞下一个连接
		//如果 使用的话 每一个连接之间是并发的 可以连接多个客户端
	}
}

func handleConn(conn net.Conn, pos *time.Location) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().In(pos).Format("15:04:05")+"--in--"+pos.String()+"\n")
		if err != nil {
			return
		}
		wait := rand.New(rand.NewSource(99))
		time.Sleep(time.Duration(wait.Intn(5)) * time.Second)
	}

}
