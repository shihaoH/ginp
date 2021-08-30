package ginp

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

type Gp struct {
	*gin.Engine
	group *gin.RouterGroup
	props []interface{}
}

// Construct 构造函数
func Construct() *Gp {
	g := &Gp{
		Engine: gin.New(),
		props: make([]interface{}, 0),
	}
	g.Use(ErrHandler())
	return g
}

// Launch 启动，默认使用8080端口
func (g *Gp) Launch(addr ...string) {
	if len(addr) > 0 {
		g.Run(addr...)
	} else {
		g.Run(":8080")
	}
}

// Mount 挂载，将实现出的接口挂载到gp
func (g *Gp) Mount(group string, control ...Controller) *Gp {
	g.group = g.Group(group)
	for _, c := range control {
		c.Build(g)
		g.setProp(c)
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
func (g *Gp) Beans(prop ...interface{}) *Gp {
	g.props = append(g.props, prop...)
	return g
}

// setProp 设置属性
func (g *Gp) setProp(c Controller) {
	cla := reflect.ValueOf(c).Elem()
	for i := 0; i < cla.NumField(); i++ {
		f := cla.Field(i)
		if !f.IsNil() || f.Kind() != reflect.Ptr {
			continue
		}
		if p := g.getProp(f.Type()); p != nil {
			f.Set(reflect.New(f.Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}

// getProp 获取参数
func (g *Gp) getProp(t reflect.Type) interface{} {
	for _, p := range g.props {
		if t == reflect.TypeOf(p) {
			return p
		}
	}
	return nil
}

