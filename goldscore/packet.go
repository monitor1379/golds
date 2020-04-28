package goldscore

import (
	"bytes"
	"fmt"
	"strconv"
)

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-26 16:46:09
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-27 23:18:56
 */

type Packet struct {
	PacketType PacketType
	Value      []byte
	Array      []*Packet
}

var (
	EmptyPacket = &Packet{PacketType: PacketTypeString, Value: []byte("empty packet")}
	OKPacket    = &Packet{PacketType: PacketTypeString, Value: []byte("OK")}
)

const (
	packetStringTpl = "{ PacketType: %+v, Value(Decoded): %+v,  Array: %+v }"
)

func NewStringPacket(msg string) *Packet {
	packet := new(Packet)
	packet.PacketType = PacketTypeString
	packet.Value = []byte(msg)
	return packet
}

func NewErrorPacket(msg string) *Packet {
	buf := bytes.Buffer{}
	_, err := buf.WriteString("Error: ")
	if err != nil {
		return nil
	}
	_, err = buf.WriteString(msg)
	if err != nil {
		return nil
	}

	packet := new(Packet)
	packet.PacketType = PacketTypeError
	packet.Value = buf.Bytes()
	return packet
}

func NewIntPacket(i int) *Packet {
	packet := new(Packet)
	packet.PacketType = PacketTypeInt
	packet.Value = []byte(strconv.Itoa(i))
	return packet
}

func NewBulkStringPacket(data []byte) *Packet {
	packet := new(Packet)
	packet.PacketType = PacketTypeBulkString
	packet.Value = data
	return packet
}

func NewEmptyArrayPacket() *Packet {
	packet := new(Packet)
	packet.PacketType = PacketTypeArray
	return packet
}

func (this *Packet) Add(subPacket *Packet) *Packet {
	this.Array = append(this.Array, subPacket)
	return this
}

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
