package goldscore_test

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-26 16:44:03
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-27 23:25:02
 */

import (
	"fmt"
	"testing"

	"github.com/monitor1379/golds/goldscore"
)

func TestPacketType(t *testing.T) {
	fmt.Println(goldscore.PacketTypeString)
	fmt.Println(goldscore.PacketTypeError)
	fmt.Println(goldscore.PacketTypeInt)
	fmt.Println(goldscore.PacketTypeBulkString)
	fmt.Println(goldscore.PacketTypeArray)

	pt := goldscore.PacketType(' ')
	fmt.Println(pt)

}
