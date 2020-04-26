/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-26 16:47:46
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-26 19:40:23
 */

package golds

import (
	"bufio"
	"io"
)

type PacketDecoder struct {
	reader *bufio.Reader
}

func NewPacketDecoder(reader io.Reader) *PacketDecoder {
	return &PacketDecoder{
		reader: bufio.NewReader(reader),
	}
}

func (this *PacketDecoder) Decode() (*Packet, error) {
	return this.decode()
}

func (this *PacketDecoder) decode() (*Packet, error) {
	firstByte, err := this.reader.ReadByte()
	if err != nil {
		return nil, err
	}

	packet := new(Packet)
	packet.PacketType = PacketType(firstByte)

	switch packet.PacketType {
	case PacketTypeString, PacketTypeError, PacketTypeInt:
		packet.Value, err = this.decodeBytes()
	case PacketTypeBulkString:
		packet.Value, err = this.decodeBulkBytes()
	case PacketTypeArray:
		packet.Array, err = this.decodeArray()
	default:
		return nil, ErrInvalidPacketType
	}

	if err != nil {
		return nil, err
	}
	return packet, nil
}

func (this *PacketDecoder) decodeBytes() ([]byte, error) {
	data, _, err := this.reader.ReadLine()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *PacketDecoder) decodeBulkBytes() ([]byte, error) { return nil, nil }
func (this *PacketDecoder) decodeArray() ([]*Packet, error)  { return nil, nil }
