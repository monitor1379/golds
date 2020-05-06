package main

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-25 13:00:06
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-26 16:33:53
 */

import (
	"flag"
	"fmt"

	"github.com/monitor1379/golds"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	path string
	port int
	help bool
)

func init() {
	flag.StringVar(&path, "path", "./db", "database directory")
	flag.IntVar(&port, "port", 1379, "port")
	flag.BoolVar(&help, "help", false, "help")
	flag.Parse()
}

func main() {
	if help {
		flag.Usage()
		return
	}

	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		panic(err)
	}

	server := golds.NewServer(db)

	err = server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
