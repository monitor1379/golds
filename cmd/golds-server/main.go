package main

import (
	"fmt"

	"github.com/monitor1379/golds"

	"github.com/syndtr/goleveldb/leveldb"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-25 13:00:06
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-25 13:08:59
 */
func main() {
	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(db)

	server := golds.NewServer()
	err = server.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
