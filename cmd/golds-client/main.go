package main

import (
	"fmt"

	"github.com/monitor1379/golds"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-25 13:00:24
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-28 22:36:04
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

	value, err := client.Get([]byte("key1"))
	if err != nil {
		panic(err)
	}
	fmt.Println("get:", string(value))
}
