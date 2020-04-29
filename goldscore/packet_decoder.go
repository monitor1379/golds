package goldscore

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-26 16:47:46
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-29 13:08:30
 */

import (
	"bufio"
	"errors"
	"io"
)

const (
	MaxBulkBytesLength = 1024 * 1024 * 512 // 512MB
	MaxArrayLength     = 1024 * 1024 * 1   // 1M
)

var (
	ErrInvalidLF = errors.New("invalid line feed")

	ErrInvalidBulkBytesLength = errors.New("invalid bulk bytes length")
	ErrTooLongBulkBytesLength = errors.New("bulk bytes length is too long")

	ErrInvalidArrayLength = errors.New("invalid array length")
	ErrTooLongArrayLength = errors.New("array length is too long")
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

	data := make([]byte, n+1)

	_, err = io.ReadFull(this.reader, data)
	if err != nil {
		return nil, err
	}

	if data[n] != '\n' {
		return nil, ErrInvalidLF
	}

	return data[:n], nil
}

func (this *PacketDecoder) decodeArray() ([]*Packet, error) {
	n, err := this.decodeInt()
	if err != nil {
		return nil, err
	}

	if n < 0 {
		return nil, ErrInvalidArrayLength
	}

	if n == 0 {
		return []*Packet{}, nil
	}

	if n > MaxArrayLength {
		return nil, ErrTooLongArrayLength
	}

	var packets []*Packet
	for i := 0; i < n; i++ {
		packet, err := this.decode()
		if err != nil {
			return nil, err
		}
		packets = append(packets, packet)
	}

	return packets, nil
}
