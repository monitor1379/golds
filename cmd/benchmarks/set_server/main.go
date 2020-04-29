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
 * @Last Modified time: 2020-04-29 13:20:15
 */

func main() {
	nClient := 100
	nKey := 1000
	valueSize := 100

	keys := make([][]byte, nKey)
	value := make([]byte, valueSize)
	fmt.Println(len(value))
	for i := 0; i < nKey; i++ {
		keys[i] = []byte(strconv.Itoa(i))
	}

	for i := 0; i < len(value); i++ {
		value[i] = []byte(strconv.Itoa(i % 10))[0]
	}

	var waitGroup sync.WaitGroup
	startTime := time.Now()
	for i := 0; i < nClient; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			client, err := golds.Dial(":3000")
			if err != nil {
				panic(err)
			}

			defer client.Close()
			for _, key := range keys {
				err = client.Set(key, value)
				if err != nil {
					fmt.Println("err:", err)
				}
			}
		}()
	}

	waitGroup.Wait()
	endTime := time.Now()
	fmt.Println("total time cost:", endTime.Sub(startTime))
}
