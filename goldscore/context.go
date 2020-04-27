package goldscore

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-27 23:07:42
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-28 00:51:29
 */
type Context struct {
	kvStorage map[string]interface{}
}

func NewContext() *Context {
	ctx := new(Context)
	ctx.kvStorage = make(map[string]interface{})
	return ctx
}

func (this *Context) Set(k string, v interface{}) {
	this.kvStorage[k] = v
}

func (this *Context) Get(k string) (interface{}, bool) {
	v, ok := this.kvStorage[k]
	return v, ok
}
