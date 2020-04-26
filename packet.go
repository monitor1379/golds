package golds

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-26 16:46:09
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-26 16:46:43
 */

type Packet struct {
	PacketType PacketType
	Value      []byte
	Array      []*Packet
}
