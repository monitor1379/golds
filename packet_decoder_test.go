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
