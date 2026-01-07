package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// 定义变量 -> 定义参数 -> 有默认值
	host := flag.String("host", "127.0.0.1", "host to check")
	prot := flag.Int("port", 80, "port to check")
	timeout := flag.Duration("timeout", 3*time.Second, "timeout")

	// 解析
	flag.Parse()

	// port check
	if *prot <= 0 || *prot >= 65535 {
		fmt.Println("port must be between 1 and 65535")
		os.Exit(1)
	}

	// 格式化
	addr := fmt.Sprintf("%s:%d", *host, *prot)

	// 端口探测
	conn, err := net.DialTimeout("tcp", addr, *timeout)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error connecting", addr, err)
		os.Exit(1)
	}

	// 断开连接
	conn.Close()

	// 结果输出
	fmt.Printf("Checking %s with timeout %s \n", addr, *timeout)
	os.Exit(0)
}
