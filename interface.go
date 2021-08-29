package ginp

import "github.com/gin-gonic/gin"

// Controller 控制器接口，实现Build以供挂载使用
type Controller interface {
	Build(gp *Gp)
}

// Fairing 中间件函数接口，提供编写中间件函数
type Fairing interface {
	Func(*gin.Context) error
}
