package main

import (
	"github.com/monitor1379/golds"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-25 13:00:24
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-27 22:52:03
 */

// TODO(monitor1379): 服务器需要实现心跳检测

func main() {
	client, err := golds.Dial(":3000")
	if err != nil {
		panic(err)
	}

	defer client.Close()

	err = client.Set([]byte("key1"), []byte("value1"))
	if err != nil {
		panic(err)
	}
}
