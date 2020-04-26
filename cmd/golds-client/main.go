package main

import "net"

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-25 13:00:24
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-26 16:35:03
 */

func main() {
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
}
