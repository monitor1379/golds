package golds

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-26 16:41:01
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-26 16:43:49
 */

import (
	"errors"
	"fmt"
)

type PacketType byte

var (
	PacketTypeString     PacketType = '+'
	PacketTypeError      PacketType = '-'
	PacketTypeInt        PacketType = ':'
	PacketTypeBulkString PacketType = '$'
	PacketTypeArray      PacketType = '*'

	ErrInvalidPacketType = errors.New("invalid packet type")
)

func (pt PacketType) String() string {
	var s string
	switch pt {
	case PacketTypeString:
		s = "string"
	case PacketTypeError:
		s = "error"
	case PacketTypeInt:
		s = "int"
	case PacketTypeBulkString:
		s = "bulk string"
	case PacketTypeArray:
		s = "array"
	default:
		s = "unknown type"
	}
	return fmt.Sprintf("{ '%v': (%s) }", string(byte(pt)), s)
}
