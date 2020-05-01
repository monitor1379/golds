package handlers

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-01 19:32:19
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-01 19:37:11
 */

import (
	"fmt"

	"github.com/monitor1379/golds/goldscore"
)

func Delete(ctx *goldscore.Context) {
	requestPacket := ctx.GetRequestPacket()
	if len(requestPacket.Array) != 2 {
		ctx.SetRequestPacket(goldscore.NewErrorPacket("invalid packet format for command 'delete'"))
		return
	}

	err := ctx.DB().Delete(requestPacket.Array[1].Value, nil)
	if err != nil {
		ctx.SetResponsePacket(goldscore.NewErrorPacket(fmt.Sprintf("delete record failed")))
		return
	}

	ctx.SetResponsePacket(goldscore.OKPacket)
}
