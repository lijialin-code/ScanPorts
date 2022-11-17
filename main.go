package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"time"
)

func tcpGather(ip string, ports []string) map[string]string {

	results := make(map[string]string)
	for _, port := range ports {
		address := net.JoinHostPort(ip, port)
		// 3 秒超时
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			results[port] = "failed"
			// todo log handler
		} else {
			if conn != nil {
				results[port] = "success"
				_ = conn.Close()
			} else {
				results[port] = "failed"
			}
		}
	}
	return results
}

func main() {
	var (
		ip    string
		ports string
	)
	flag.StringVar(&ip, "ip", "127.0.0.1", "ip address")
	flag.StringVar(&ports, "ports", "22,80", "port list")
	flag.Parse()
	slicePort := strings.Split(ports, ",")
	res := tcpGather(ip, slicePort)
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(res, now)
}
