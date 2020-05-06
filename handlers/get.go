/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-27 23:26:29
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-27 23:26:49
 */
package handlers

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/monitor1379/golds/goldscore"
)

const (
	CommandNameGet = "get"
)

func Get(ctx *goldscore.Context) {
	requestPacket := ctx.GetRequestPacket()
	if len(requestPacket.Array) != 2 {
		ctx.SetResponsePacket(goldscore.NewErrorPacket("invalid packet format for command 'get'"))
		return
	}

	value, err := ctx.DB().Get(requestPacket.Array[1].Value, nil)

	if err == leveldb.ErrNotFound {
		ctx.SetResponsePacket(goldscore.NewBulkStringPacket(nil))
		return
	}

	if err != nil {
		ctx.SetResponsePacket(goldscore.NewErrorPacket(fmt.Sprintf("read db failed(%s)", err)))
		return
	}

	ctx.SetResponsePacket(goldscore.NewBulkStringPacket(value))
}
