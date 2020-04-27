package golds_test

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-26 16:47:11
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-26 19:31:37
 */

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/monitor1379/golds"
)

func TestPacketDecoder(t *testing.T) {
	packetStrings := []string{
		"+OK\n",
		"-Error message\n",
		"$5\nhello\n",
		"$11\nhello\nworld\n",
		"$0\n\n",
		"$-1\n",
		"*1\n:123\n",
		"*3\n$3\nSET\n$3\nkey\n$5\nvalue\n",
		"*0\n",
	}

	for _, packetString := range packetStrings {
		packet, err := golds.NewPacketDecoder(bytes.NewBufferString(packetString)).Decode()
		if err != nil {
			t.Errorf("ERROR: decode error: %s", err)
			continue
		}
		fmt.Println("packet:", packet)
	}

}
