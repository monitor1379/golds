/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-29 11:59:51
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-29 12:05:52
 */

package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		panic(err)
	}

	n := 10000
	valueSize := 1024

	keys := make([][]byte, n)
	value := make([]byte, valueSize)

	for i := 0; i < n; i++ {
		keys[i] = []byte(strconv.Itoa(i))
	}

	for i := 0; i < len(value); i++ {
		value[i] = []byte(strconv.Itoa(i % 10))[0]
	}

	startTime := time.Now()
	for i := 0; i < n; i++ {
		db.Put(keys[i], value, nil)
	}
	endTime := time.Now()
	fmt.Println("total time cost:", endTime.Sub(startTime))
}
