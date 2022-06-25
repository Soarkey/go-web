package engine

import (
	"net/http"
)

// HandleFunc 定义请求的 handler
type HandleFunc func(*Context)

// Engine 实现 ServeHTTP 接口
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

// ServeHTTP 主入口
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}

func (e *Engine) addRoute(method, pattern string, handler HandleFunc) {
	e.router.addRoute(method, pattern, handler)
}

// GET 添加 GET 请求路由
func (e *Engine) GET(pattern string, handler HandleFunc) {
	e.addRoute("GET", pattern, handler)
}

// POST 添加 POST 请求路由
func (e *Engine) POST(pattern string, handler HandleFunc) {
	e.addRoute("POST", pattern, handler)
}

// Run 启动 http server
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

// New 构造函数
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return &Engine{router: newRouter()}
}

// RouterGroup 路由分组
type RouterGroup struct {
	prefix      string       // 前缀
	middlewares []HandleFunc // 中间件
	parent      *RouterGroup // 支持组合
	engine      *Engine      // 所有分组共享一个Engine
}
