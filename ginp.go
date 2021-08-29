package ginp

import (
	"github.com/gin-gonic/gin"
)

type Gp struct {
	*gin.Engine
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
func (g *Gp) Mount(is ...Interface) *Gp {
	for _, i := range is {
		i.Build(g)
	}
	return g
}

