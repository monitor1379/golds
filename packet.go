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

	switch this.PacketType {
	case PacketTypeString:
		return fmt.Sprintf(packetStringTpl, this.PacketType, strconv.Quote(this.MustGetValueAsString()), this.Array)
	case PacketTypeError:
		return fmt.Sprintf(packetStringTpl, this.PacketType, strconv.Quote(this.MustGetValueAsString()), this.Array)
	case PacketTypeInt:
		return fmt.Sprintf(packetStringTpl, this.PacketType, this.MustGetValueAsInt(), this.Array)
	case PacketTypeBulkString:
		return fmt.Sprintf(packetStringTpl, this.PacketType, strconv.Quote(this.MustGetValueAsString()), this.Array)
	case PacketTypeArray:
		return fmt.Sprintf(packetStringTpl, this.PacketType, strconv.Quote(this.MustGetValueAsString()), this.Array)
	default:
		return ""
	}

}
