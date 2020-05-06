package goldscore

import "github.com/syndtr/goleveldb/leveldb"

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-27 23:07:42
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-28 00:51:29
 */
type Context struct {
	db             *leveldb.DB
	requestPacket  *Packet
	responsePacket *Packet

	kvStorage map[string]interface{}
}

func NewContext(db *leveldb.DB, requestPacket *Packet) *Context {
	ctx := new(Context)
	ctx.db = db
	ctx.requestPacket = requestPacket
	ctx.kvStorage = make(map[string]interface{})
	return ctx
}

func (this *Context) DB() *leveldb.DB {
	return this.db
}

func (this *Context) SetRequestPacket(requestPacket *Packet) {
	this.requestPacket = requestPacket
}

func (this *Context) GetRequestPacket() *Packet {
	return this.requestPacket
}

func (this *Context) SetResponsePacket(responsePacket *Packet) {
	this.responsePacket = responsePacket
}

func (this *Context) GetResponsePacket() *Packet {
	return this.responsePacket
}

func (this *Context) Set(k string, v interface{}) {
	this.kvStorage[k] = v
}

func (this *Context) Get(k string) (interface{}, bool) {
	v, ok := this.kvStorage[k]
	return v, ok
}
