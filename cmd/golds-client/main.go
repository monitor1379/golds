package main

import (
	"flag"
	"fmt"

	"github.com/monitor1379/golds"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-25 13:00:24
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-01 19:56:27
 */

var (
	host string
	port int
	help bool
)

func init() {
	flag.StringVar(&host, "host", "localhost", "server host")
	flag.IntVar(&port, "port", 1379, "server port")
	flag.BoolVar(&help, "help", false, "help")
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	client, err := golds.Dial(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Printf("ERROR: Connect server failed. error = '%s'.\n", err)
		return
	}
	fmt.Println(client)

}
