package main

import (
	"fmt"

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
	// test1()
	test2()
}

func test1() {

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

func test2() {
	client, err := golds.Dial(":3000")
	if err != nil {
		panic(err)
	}

	defer client.Close()

	client.Set([]byte("key1"), []byte("value1"))
	value, _ := client.Get([]byte("key1"))
	fmt.Println("get:", string(value))

	client.Set([]byte("key2"), []byte("value2"))
	value, _ = client.Get([]byte("key2"))
	fmt.Println("get:", string(value))

	client.Set([]byte("key1"), []byte("value3"))
	value, _ = client.Get([]byte("key1"))
	fmt.Println("get:", string(value))
}
