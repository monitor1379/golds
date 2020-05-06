package goldscore

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-27 23:07:14
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-28 00:16:43
 */

type Router struct {
	handleFuncsMap map[string]HandleFunc
}

func NewRouter() *Router {
	router := new(Router)
	router.handleFuncsMap = make(map[string]HandleFunc)
	return router
}

func (this *Router) AddHandleFunc(routePath string, handleFunc HandleFunc) {
	this.handleFuncsMap[routePath] = handleFunc
}

func (this *Router) Route(routePath string) (HandleFunc, bool) {
	v, ok := this.handleFuncsMap[routePath]
	return v, ok
}
