package golds

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-27 21:10:58
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-01 22:16:53
 */

import (
	"fmt"
	"net"
	"strconv"
	"sync"

	"github.com/monitor1379/golds/handlers"

	"github.com/monitor1379/golds/goldscore"
	"go.uber.org/zap"
)

type Client struct {
	mu            sync.Mutex
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
		mu:            sync.Mutex{},
		conn:          conn,
		packetEncoder: goldscore.NewPacketEncoder(conn),
		packetDecoder: goldscore.NewPacketDecoder(conn),
	}
	return client, nil
}

func (this *Client) Close() error {
	return this.conn.Close()
}

func (this *Client) do(requestPacket *goldscore.Packet) (*goldscore.Packet, error) {
	err := this.packetEncoder.Encode(requestPacket)
	if err != nil {
		return nil, err
	}

	responsePacket, err := this.packetDecoder.Decode()
	if err != nil {
		return nil, err
	}

	if responsePacket.PacketType == goldscore.PacketTypeError {
		return nil, fmt.Errorf("golds server error: %s", string(responsePacket.Value))
	}

	return responsePacket, nil
}

func (this *Client) Set(key, value []byte) error {
	this.mu.Lock()
	defer this.mu.Unlock()

	requestPacket := goldscore.NewEmptyArrayPacket().
		Add(goldscore.NewBulkStringPacket([]byte(handlers.CommandNameSet))).
		Add(goldscore.NewBulkStringPacket(key)).
		Add(goldscore.NewBulkStringPacket(value))

	responsePacket, err := this.do(requestPacket)
	if err != nil {
		return err
	}

	logger.Debug("response packet", zap.String("packet", responsePacket.String()))
	return nil
}

func (this *Client) Get(key []byte) ([]byte, error) {
	this.mu.Lock()
	defer this.mu.Unlock()

	requestPacket := goldscore.NewEmptyArrayPacket().
		Add(goldscore.NewBulkStringPacket([]byte(handlers.CommandNameGet))).
		Add(goldscore.NewBulkStringPacket(key))

	responsePacket, err := this.do(requestPacket)
	if err != nil {
		return nil, err
	}

	return responsePacket.Value, nil
}

func (this *Client) Del(key []byte) error {
	this.mu.Lock()
	defer this.mu.Unlock()

	requestPacket := goldscore.NewEmptyArrayPacket().
		Add(goldscore.NewBulkStringPacket([]byte(handlers.CommandNameDel))).
		Add(goldscore.NewBulkStringPacket(key))

	responsePacket, err := this.do(requestPacket)
	if err != nil {
		return err
	}
	_ = responsePacket

	return nil
}

func (this *Client) Keys() ([]byte, error) {
	this.mu.Lock()
	defer this.mu.Unlock()

	requestPacket := goldscore.NewEmptyArrayPacket().
		Add(goldscore.NewBulkStringPacket([]byte(handlers.CommandNameKeys)))

	responsePacket, err := this.do(requestPacket)
	if err != nil {
		return nil, err
	}

	for i, subPacket := range responsePacket.Array {
		fmt.Printf("%d): %s\n", i, strconv.Quote(string(subPacket.Value)))
	}

	return nil, nil
}
