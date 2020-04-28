package handlers

import (
	"fmt"

	"github.com/monitor1379/golds/goldscore"
)

func Set(ctx *goldscore.Context) {
	requestPacket := ctx.GetRequestPacket()
	if len(requestPacket.Array) != 3 {
		ctx.SetResponsePacket(goldscore.NewErrorPacket("invalid packet format for command 'set'"))
		return
	}

	err := ctx.DB().Put(requestPacket.Array[1].Value, requestPacket.Array[2].Value, nil)
	if err != nil {
		ctx.SetResponsePacket(goldscore.NewErrorPacket(fmt.Sprintf("write db failed(%s)", err)))
		return
	}

	ctx.SetResponsePacket(goldscore.OKPacket)
}
