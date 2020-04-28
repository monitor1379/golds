package goldscore_test

import (
	"fmt"
	"testing"

	"github.com/monitor1379/golds/goldscore"
)

func TestPacket(t *testing.T) {
	fmt.Println(goldscore.NewStringPacket(""))
	fmt.Println(goldscore.NewStringPacket("hello world"))
	fmt.Println(goldscore.NewStringPacket("OK"))
	fmt.Println(goldscore.NewStringPacket("PING"))

	fmt.Println(goldscore.NewErrorPacket(""))
	fmt.Println(goldscore.NewErrorPacket("error message"))

	fmt.Println(goldscore.NewIntPacket(0))
	fmt.Println(goldscore.NewIntPacket(12345678))
	fmt.Println(goldscore.NewIntPacket(-1))
	fmt.Println(goldscore.NewIntPacket(-12345678))

	fmt.Println(goldscore.NewBulkStringPacket([]byte("")))
	fmt.Println(goldscore.NewBulkStringPacket([]byte("hello")))
	fmt.Println(goldscore.NewBulkStringPacket([]byte("hello world")))
	fmt.Println(goldscore.NewBulkStringPacket([]byte("hello\nworld")))
	fmt.Println(goldscore.NewBulkStringPacket([]byte("hello\nworld\n")))
	fmt.Println(goldscore.NewBulkStringPacket(nil))

	fmt.Println(goldscore.NewEmptyArrayPacket().Add(goldscore.NewStringPacket("set")))
	fmt.Println(goldscore.NewEmptyArrayPacket().
		Add(goldscore.NewBulkStringPacket([]byte("set"))).
		Add(goldscore.NewBulkStringPacket([]byte("key1"))).
		Add(goldscore.NewBulkStringPacket([]byte("value1"))),
	)
}
