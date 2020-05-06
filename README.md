# Golds: Go LevelDB Service

Golds(GO LevelDb Service)是一个轻量级的LevelDB网络服务器（协议）。该项目旨在提供一个基于磁盘的、简易的、高性能的KV数据库。

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

# 通信协议

Golds的TCP通信协议采取类似Redis RESP协议的设计方式，特点为:
- 实现简单
- 快速地被计算机解析
- 人类可读


## 网络层

Golds在(默认)端口1379上监听到来的TCP连接。客户端连接到来时，Golds服务器会创建一个TCP连接。在客户端与服务器端之间传输的每个Golds命令或者数据都以`\n`结尾。


## 请求与响应

Golds接收由不同参数组成的命令。一旦收到命令，将会立刻被处理，并回复给客户端。

## 协议定义

Golds中，一个请求数据包或者一个响应数据包表示为一段字节数组。通过判断第一个字节，来判断该数据包的数据结构类型:

- `+`: 单行字符串
- `-`: 错误消息
- `:`: 整型数字
- `$`: 二进制安全字符串(Bulk String)，也可以理解为字节串
- `*`: 数组

例如:

- `+OK\n`: 该字节流会被解析成一个单行字符串数据包
- `-Error: read db failed\n`: 会被解析成一个错误消息数据包
- `:123\n`: 表示整数`123`
- `$11\nhello\nworld\n`: 表示一个长度为11的字符串`hello\nworld`
- `*2\n:123\n$5\nhello\n`: `*2`表示接下来的是一个数组，长度为2, 第一个元素为`:123`，整数类型，内容为123。第二个元素为`$5\nhello`，一个长度为5个字节的字节串，内容为`hello`

通过以上基础数据结构的序列化定义，我们可以设计命令的序列化方式。

例如，`set key1 value1`命令的序列化数据为:
```
*3
$3
set
$4
key1
$5
value1
```


`keys`命令的序列化数据为:
```
*1
$4
keys
```

我们也可以通过Linux平台下的nc工具进行测试。例如，启动了server后，我们通过:
```bash
$ nc localhost 1379
```
来连接客户端。然后输入:
```
*1
$4
keys
```
如果数据库中只存在两个key，分别为`k1`和`k2`，那么可以看到控制台中会打印:
```
*2
$2
k1
$2
k2
```

具体的代码可以参考`github.com/monitor1379/golds/goldscore`里的:
- packet.go: 定义了数据包Packet的格式
- packet_type.go: 定义了数据包的类型
- packet_decoder.go: 从二进制数据流中解析出Packet的反序列化方法
- packet_encoder.go: 将Packet序列化为二进制数据流的序列化方法
