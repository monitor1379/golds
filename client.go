package golds

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-27 21:10:58
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-27 23:22:41
 */

import (
	"errors"
	"net"

	"github.com/monitor1379/golds/goldscore"
)

type Client struct {
	conn          net.Conn
	packetEncoder *goldscore.PacketEncoder
	packetDecoder *goldscore.PacketDecoder
}

func Dial(address string) (*Client, error) {
	conn, err := net.Dial(defaultNetwork, address)
	if err != nil {
		return nil, err
	}
	client := &Client{
		conn:          conn,
		packetEncoder: goldscore.NewPacketEncoder(conn),
		packetDecoder: goldscore.NewPacketDecoder(conn),
	}
	return client, nil
}

func (this *Client) Close() error {
	return this.conn.Close()
}

func (this *Client) Set(key, value []byte) error {
	requestPacket := goldscore.Packet{
		PacketType: goldscore.PacketTypeArray,
		Array: []*goldscore.Packet{
			&goldscore.Packet{goldscore.PacketType: goldscore.PacketTypeBulkString, Value: []byte("set")},
			&goldscore.Packet{goldscore.PacketType: goldscore.PacketTypeBulkString, Value: key},
			&goldscore.Packet{goldscore.PacketType: goldscore.PacketTypeBulkString, Value: value},
		},
	}

	err := this.packetEncoder.Encode(&requestPacket)
	if err != nil {
		return err
	}

	responsePacket, err := this.packetDecoder.Decode()
	if err != nil {
		return err
	}

	if responsePacket.PacketType == goldscore.PacketTypeError {
		return errors.New(string(responsePacket.Value))
	}

	return nil
}
