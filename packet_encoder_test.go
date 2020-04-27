package golds_test

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-26 16:47:11
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-26 16:47:31
 */

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/monitor1379/golds"
)

func TestPacketEncoder(t *testing.T) {
	packets := []*golds.Packet{
		&golds.Packet{PacketType: golds.PacketTypeString, Value: []byte("hello")},
		&golds.Packet{PacketType: golds.PacketTypeString, Value: []byte("hello world")},
		&golds.Packet{PacketType: golds.PacketTypeString, Value: []byte("")},
		&golds.Packet{PacketType: golds.PacketTypeError, Value: []byte("Error message")},
		&golds.Packet{PacketType: golds.PacketTypeBulkString, Value: []byte("hello")},
		&golds.Packet{PacketType: golds.PacketTypeBulkString, Value: []byte("hello world")},
		&golds.Packet{PacketType: golds.PacketTypeBulkString, Value: []byte("")},
		&golds.Packet{PacketType: golds.PacketTypeBulkString, Value: nil},
		&golds.Packet{PacketType: golds.PacketTypeBulkString, Value: []byte("hello\nworld")},
		&golds.Packet{PacketType: golds.PacketTypeBulkString, Value: []byte("hello\nworld\n")},
	}

	for _, packet := range packets {
		buf := bytes.Buffer{}
		packetEncoder := golds.NewPacketEncoder(&buf)
		err := packetEncoder.Encode(packet)
		if err != nil {
			t.Errorf("ERROR: encode error: %s", err)
			continue
		}
		fmt.Println(strconv.Quote(buf.String()))
	}
}
