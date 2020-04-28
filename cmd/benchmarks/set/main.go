package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/monitor1379/golds"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-28 22:35:10
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-28 23:35:05
 */

func main() {
	client, err := golds.Dial(":3000")
	if err != nil {
		panic(err)
	}

	defer client.Close()

	n := 1024
	keys := make([][]byte, n)
	value := make([]byte, 1024)
	fmt.Println(len(value))
	for i := 0; i < n; i++ {
		keys[i] = []byte(strconv.Itoa(i))
	}

	for i := 0; i < len(value); i++ {
		value[i] = []byte(strconv.Itoa(i % 10))[0]
	}

	var waitGroup sync.WaitGroup
	startTime := time.Now()
	for _, key := range keys {
		waitGroup.Add(1)
		go func(key []byte) {
			err = client.Set(key, value)
			if err != nil {
				fmt.Println("err:", err)
			}
			waitGroup.Done()
		}(key)

	}
	waitGroup.Wait()
	endTime := time.Now()
	fmt.Println("total time cost:", endTime.Sub(startTime))
}
