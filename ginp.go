package ginp

import (
	"github.com/gin-gonic/gin"
)

type Gp struct {
	*gin.Engine
	group *gin.RouterGroup
}

// Construct 构造函数
func Construct() *Gp {
	return &Gp{
		Engine: gin.New(),
	}
}

// Launch 启动，默认使用8080端口
func (g *Gp) Launch(addr ...string) {
	if len(addr) > 0 {
		g.Run(addr...)
	} else {
		g.Run(":8080")
	}
}

// Mount 挂载，将实现出的接口挂载到gp，链式函数
func (g *Gp) Mount(group string, is ...Interface) *Gp {
	g.group = g.Group(group)
	for _, i := range is {
		i.Build(g)
	}
	return g
}

// Handle 重写gin.Group.Handle，保证前后一致
func (g *Gp) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *Gp {
	// 可变参数传入后是切片，需要[...]解构成多参数
	g.group.Handle(httpMethod, relativePath, handlers...)
	return g
}
