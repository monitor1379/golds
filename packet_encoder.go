/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-26 16:47:46
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-26 16:51:59
 */

package golds

type PacketEncoder struct{}

func NewPacketEncoder() *PacketEncoder {
	packetEncoder := new(PacketEncoder)
	return packetEncoder
}
