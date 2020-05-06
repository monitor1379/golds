# GOLDS: Go LevelDB Service

GOLDS(GO LevelDb Service)是一个轻量级的LevelDB网络服务器（协议）。该项目旨在提供一个基于磁盘的、简易的、高性能的KV数据库。

# 功能特性

- 性能高效: 底层基于Google高性能磁盘KV数据库LevelDB。
- 协议简单: 采用类似Redis通信协议RESP的TCP数据传输协议，解析快速，人类可读。

# 源码安装

```bash
$ go get github.com/monitor1379/golds
$ cd $GOPATH/github.com/monitor1379/golds/cmd/golds-server
$ go build

$ cd $GOPATH/github.com/monitor1379/golds/cmd/golds-client
$ go build
```

# 使用

启动server的方式以及其默认参数:
```bash
$ golds-server
$ golds-server --path=./db --port=1379
```


启动client的方式以及其默认参数:
```bash
$ golds-client
$ golds-client --host=localhost --port=1379
```


client中的命令使用类似redis:
```bash
(localhost:1379):SET k1 v1
OK

(localhost:1379):GET k1
1): "v1"

(localhost:1379):set k2 v2
OK

(localhost:1379):get k2
1): "v2"

(localhost:1379):KEYS
0): "k1"
1): "k2"
```

# API

主要接口:
- `func Dial(address string) (*golds.Client, error)`: 连接服务器
- `func (*golds.Client) Set(key, value []byte) error`: SET操作
- `func (*golds.Client) Get(key []byte) ([]byte, error)`: GET操作
- `func (*golds.Client) Keys([][]byte, error)`: KEYS操作


示例代码:
```go
package main

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

```


# 基准测试