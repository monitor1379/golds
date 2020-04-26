/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-26 16:47:46
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-26 19:40:23
 */

package golds

import (
	"bufio"
	"errors"
	"io"
)

const (
	MaxBulkBytesLength = 1024 * 1024 * 1 // 1MB
	MaxArrayLength     = 1024 * 1024 * 1 // 1M
)

var (
	ErrInvalidBulkBytesLength = errors.New("invalid bulk bytes length")
	ErrTooLongBulkBytesLength = errors.New("bulk bytes length is too long")
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

func (this *PacketDecoder) decodeInt() (int, error) {
	data, _, err := this.reader.ReadLine()
	if err != nil {
		return 0, err
	}

	n, err := Btoi64(data)
	if err != nil {
		return 0, err
	}

	return int(n), nil
}

func (this *PacketDecoder) decodeBulkBytes() ([]byte, error) {
	n, err := this.decodeInt()
	if err != nil {
		return nil, err
	}

	if n < -1 {
		return nil, ErrInvalidBulkBytesLength
	}

	if n == -1 {
		return nil, nil
	}

	if n > MaxBulkBytesLength {
		return nil, ErrTooLongBulkBytesLength
	}

	data := make([]byte, n)

	_, err = this.reader.Read(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (this *PacketDecoder) decodeArray() ([]*Packet, error) { return nil, nil }
