package main

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-25 13:00:06
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-26 16:33:53
 */

import (
	"github.com/monitor1379/golds"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		panic(err)
	}

	server := golds.NewServer(db)

	err = server.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
