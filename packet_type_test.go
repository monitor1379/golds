package golds_test

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-26 16:44:03
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-26 16:45:18
 */

import (
	"fmt"
	"testing"

	"github.com/monitor1379/golds"
)

func TestPacketType(t *testing.T) {
	fmt.Println(golds.PacketTypeString)
	fmt.Println(golds.PacketTypeError)
	fmt.Println(golds.PacketTypeInt)
	fmt.Println(golds.PacketTypeBulkString)
	fmt.Println(golds.PacketTypeArray)

	pt := golds.PacketType(' ')
	fmt.Println(pt)

}
