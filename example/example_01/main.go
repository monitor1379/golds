package main

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-06 11:35:16
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-06 11:42:42
 */

import (
	"fmt"

	"github.com/monitor1379/golds"
)

func main() {
	client, err := golds.Dial("localhost:1379")
	defer client.Close()

	if err != nil {
		panic(err)
	}

	// set k1 v1
	err = client.Set([]byte("k1"), []byte("v1"))
	if err != nil {
		panic(err)
	}

	// get k1
	value, err := client.Get([]byte("k1"))
	if err != nil {
		panic(err)
	}
	fmt.Println("value of 'k1' = ", string(value))

	// keys
	keys, err := client.Keys()
	if err != nil {
		panic(err)
	}

	for _, key := range keys {
		fmt.Println(string(key))
	}

	// del k1
	err = client.Del([]byte("k1"))
	if err != nil {
		panic(err)
	}

	// get k1
	value, err = client.Get([]byte("k1"))
	if value != nil {
		fmt.Println("Error: value should be nil because db does not have key 'k1'")
		return
	}

}
