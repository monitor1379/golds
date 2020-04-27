/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-26 16:47:46
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-26 16:51:59
 */

package golds

import (
	"bufio"
	"io"
	"strconv"
)

type PacketEncoder struct {
	writer *bufio.Writer
}

func NewPacketEncoder(writer io.Writer) *PacketEncoder {
	return &PacketEncoder{
		writer: bufio.NewWriter(writer),
	}
}

func (this *PacketEncoder) Encode(packet *Packet) error {
	return this.encode(packet)
}

func (this *PacketEncoder) encode(packet *Packet) error {
	var err error
	switch packet.PacketType {
	case PacketTypeString, PacketTypeError, PacketTypeInt:
		err = this.encodeBytes(packet)
	case PacketTypeBulkString:
		err = this.encodeBulkBytes(packet)
	case PacketTypeArray:
		err = this.encodeArray(packet)
	default:
		return ErrInvalidPacketType
	}

	if err != nil {
		return err
	}
	return nil
}

func (this *PacketEncoder) encodeBytes(packet *Packet) error {
	defer this.writer.Flush()

	err := this.writer.WriteByte(byte(packet.PacketType))
	if err != nil {
		return err
	}

	_, err = this.writer.Write(packet.Value)
	if err != nil {
		return err
	}

	err = this.writer.WriteByte('\n')
	if err != nil {
		return err
	}

	return nil
}

func (this *PacketEncoder) encodeBulkBytes(packet *Packet) error {
	defer this.writer.Flush()

	err := this.writer.WriteByte(byte(packet.PacketType))
	if err != nil {
		return err
	}

	if packet.Value == nil {
		_, err = this.writer.WriteString("-1")
		if err != nil {
			return err
		}

		err = this.writer.WriteByte('\n')
		if err != nil {
			return err
		}
		return nil
	}

	_, err = this.writer.WriteString(strconv.Itoa(len(packet.Value)))
	if err != nil {
		return err
	}

	err = this.writer.WriteByte('\n')
	if err != nil {
		return err
	}

	_, err = this.writer.Write(packet.Value)
	if err != nil {
		return err
	}

	err = this.writer.WriteByte('\n')
	if err != nil {
		return err
	}

	return nil
}

func (this *PacketEncoder) encodeArray(packet *Packet) error {
	defer this.writer.Flush()
	return nil
}
