package goldscore_test

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-26 16:47:11
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-27 23:21:35
 */

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/monitor1379/golds/goldscore"
)

func TestPacketEncoder(t *testing.T) {
	packets := []goldscore.Packet{
		goldscore.Packet{PacketType: goldscore.PacketTypeString, Value: []byte("hello")},
		goldscore.Packet{PacketType: goldscore.PacketTypeString, Value: []byte("hello world")},
		goldscore.Packet{PacketType: goldscore.PacketTypeString, Value: []byte("")},
		goldscore.Packet{PacketType: goldscore.PacketTypeError, Value: []byte("Error message")},
		goldscore.Packet{PacketType: goldscore.PacketTypeBulkString, Value: []byte("hello")},
		goldscore.Packet{PacketType: goldscore.PacketTypeBulkString, Value: []byte("hello\nworld")},
		goldscore.Packet{PacketType: goldscore.PacketTypeBulkString, Value: []byte("hello\nworld\n")},
		goldscore.Packet{PacketType: goldscore.PacketTypeBulkString, Value: []byte("hello world")},
		goldscore.Packet{PacketType: goldscore.PacketTypeBulkString, Value: []byte("")},
		goldscore.Packet{PacketType: goldscore.PacketTypeBulkString, Value: nil},
		goldscore.Packet{PacketType: goldscore.PacketTypeArray, Array: []*goldscore.Packet{
			&goldscore.Packet{PacketType: goldscore.PacketTypeBulkString, Value: []byte("SET")},
			&goldscore.Packet{PacketType: goldscore.PacketTypeBulkString, Value: []byte("key1")},
			&goldscore.Packet{PacketType: goldscore.PacketTypeBulkString, Value: []byte("value1")},
		}},

		goldscore.Packet{PacketType: goldscore.PacketTypeArray, Array: []*goldscore.Packet{
			&goldscore.Packet{PacketType: goldscore.PacketTypeInt, Value: []byte("1")},
		}},
	}

	for _, packet := range packets {
		buf := bytes.Buffer{}
		packetEncoder := goldscore.NewPacketEncoder(&buf)
		err := packetEncoder.Encode(&packet)
		if err != nil {
			t.Errorf("ERROR: encode error: %s", err)
			continue
		}
		fmt.Println(strconv.Quote(buf.String()))
	}
}
