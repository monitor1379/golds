package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/monitor1379/golds"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-25 13:00:24
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-01 20:30:46
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
	fmt.Println(client)

	reader := bufio.NewReader(os.Stdin)
	for {
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

	if commandName == "set" {
		if nCommandItems != 3 {

		}
	} else if commandName == "get" {
		if nCommandItems != 2 {

		}
	} else if commandName == "delete" {
		if nCommandItems != 2 {

		}
	} else {

	}

}
