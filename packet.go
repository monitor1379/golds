package golds

import (
	"fmt"
	"strconv"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-26 16:46:09
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-26 19:40:53
 */

type Packet struct {
	PacketType PacketType
	Value      []byte
	Array      []*Packet
}

const (
	packetStringTpl = "{ PacketType: %+v, Value(Decoded): %+v,  Array: %+v }"
)

func (this Packet) MustGetValueAsString() string {
	return string(this.Value)
}

func (this Packet) MustGetValueAsInt() int {
	i, err := Btoi64(this.Value)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func (this Packet) String() string {
	var valueStr string
	if this.Value == nil {
		valueStr = fmt.Sprintf("%+v", nil)
	} else {
		switch this.PacketType {
		case PacketTypeString, PacketTypeError, PacketTypeBulkString, PacketTypeArray:
			valueStr = strconv.Quote(this.MustGetValueAsString())
		case PacketTypeInt:
			valueStr = fmt.Sprintf("%d", this.MustGetValueAsInt())
		default:
			valueStr = fmt.Sprintf("%+v", nil)
		}
	}

	var arrayStr string
	if this.Array == nil {
		arrayStr = fmt.Sprintf("%+v", nil)
	} else {
		arrayStr = fmt.Sprintf("%+v", this.Array)
	}

	return fmt.Sprintf(packetStringTpl, this.PacketType, valueStr, arrayStr)
}
