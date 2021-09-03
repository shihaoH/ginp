package ginp

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Gp struct {
	*gin.Engine
	group       *gin.RouterGroup
	beanFactory *BeanFactory
}

// Construct 构造函数
func Construct() *Gp {
	g := &Gp{
		Engine:      gin.New(),
		beanFactory: NewBeanFactory(),
	}
	g.Use(ErrHandler())
	g.beanFactory.setBean(InitConfig())
	return g
}

// Launch 启动，默认使用8080端口
func (g *Gp) Launch() {
	var port int32 = 8080
	if conf := g.beanFactory.GetBean(new(SysConfig)); conf != nil {
		port = conf.(*SysConfig).Server.Port
	}
	g.Run(fmt.Sprintf(":%d", port))
}

// Mount 挂载，将实现出的接口挂载到gp
func (g *Gp) Mount(group string, control ...Controller) *Gp {
	g.group = g.Group(group)
	for _, c := range control {
		c.Build(g)
		g.beanFactory.inject(c)
	}
	return g
}

// Handle 封装注册
func (g *Gp) Handle(httpMethod, relativePath string, handler interface{}) *Gp {
	// 响应体转换后执行
	if h := Convert(handler); h != nil {
		g.group.Handle(httpMethod, relativePath, h)
	}
	return g
}

// Mid 中间件插入
func (g *Gp) Mid(f Fairing) *Gp {
	g.Use(func(context *gin.Context) {
		err := f.Func(context)
		if err != nil {
			// 有错误直接抛出响应
			context.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		} else {
			// 下层洋葱
			context.Next()
		}
	})
	return g
}

// Beans 依赖注入
func (g *Gp) Beans(beans ...interface{}) *Gp {
	g.beanFactory.setBean(beans)
	return g
}
