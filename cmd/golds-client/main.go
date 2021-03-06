package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/monitor1379/golds"
	"github.com/monitor1379/golds/handlers"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-25 13:00:24
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-06 11:49:28
 */

var (
	host string
	port int
	help bool
)

func init() {
	flag.StringVar(&host, "host", "localhost", "server host")
	flag.IntVar(&port, "port", 1379, "server port")
	flag.BoolVar(&help, "help", false, "help")
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	client, err := golds.Dial(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Printf("ERROR: Connect server failed. error = '%s'. \n", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\n(%s:%d):", host, port)
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Printf("ERROR: read line failed. err = '%s' .\n", err)
			continue
		}

		commandItems := strings.Split(string(line), " ")
		if len(commandItems) == 0 {
			continue
		}

		executeCommand(client, commandItems)
	}

}

func executeCommand(client *golds.Client, commandItems []string) {
	commandName := strings.ToLower(commandItems[0])
	nCommandItems := len(commandItems)

	if commandName == handlers.CommandNameSet {
		if nCommandItems != 3 {
			fmt.Printf("ERROR: invalid format for command 'set'. usage: set key value\n")
			return
		}
		err := client.Set([]byte(commandItems[1]), []byte(commandItems[2]))
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			return
		}
		fmt.Println("OK")
	} else if commandName == handlers.CommandNameGet {
		if nCommandItems != 2 {
			fmt.Printf("ERROR: invalid format for command 'get'. usage: get key\n")
			return
		}
		value, err := client.Get([]byte(commandItems[1]))
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			return
		}
		if value == nil {
			fmt.Printf("<nil>\n")
			return
		}
		fmt.Printf("1): %s\n", strconv.Quote(string(value)))
	} else if commandName == handlers.CommandNameDel {
		if nCommandItems != 2 {
			fmt.Printf("ERROR: invalid format for command 'get'. usage: get key\n")
			return
		}
		err := client.Del([]byte(commandItems[1]))
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			return
		}
		fmt.Println("OK")
	} else if commandName == handlers.CommandNameKeys {
		keys, err := client.Keys()
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			return
		}
		for i, key := range keys {
			fmt.Printf("%d): %s\n", i+1, strconv.Quote(string(key)))
		}
	} else {
		fmt.Printf("ERROR: unknown command")
	}

}
