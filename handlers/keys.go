package handlers

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-01 19:32:19
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-01 22:41:42
 */

import (
	"fmt"

	"github.com/monitor1379/golds/goldscore"
)

const (
	CommandNameKeys = "keys"
)

func Keys(ctx *goldscore.Context) {
	requestPacket := ctx.GetRequestPacket()
	if len(requestPacket.Array) != 1 {
		ctx.SetRequestPacket(goldscore.NewErrorPacket("invalid packet format for command 'keys'"))
		return
	}

	snapshot, err := ctx.DB().GetSnapshot()
	if err != nil {
		ctx.SetResponsePacket(goldscore.NewErrorPacket(fmt.Sprintf("get db snapshot failed")))
		return
	}
	defer snapshot.Release()

	iterator := snapshot.NewIterator(nil, nil)
	defer iterator.Release()

	responsePacket := goldscore.NewEmptyArrayPacket()
	for iterator.Next() {
		key := make([]byte, len(iterator.Key()))
		copy(key, iterator.Key())
		responsePacket.Add(goldscore.NewBulkStringPacket(key))
	}

	ctx.SetResponsePacket(responsePacket)
}
