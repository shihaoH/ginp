package ginp

import "github.com/gin-gonic/gin"

// ErrHandler 错误处理中间件，由Gp.Construct函数注入
func ErrHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		// recover 内置函数需在 panic 发生之后调用，所以需要使用 defer 去后置操作
		defer func() {
			if e := recover(); e != nil {
				context.AbortWithStatusJSON(400, gin.H{"error": e})
			}
		}()
		context.Next()
	}
}

// Error 错误处理函数，在需要处理error的地方调用，会走到 ErrHandler 的 recover 调用验证和统一处理
func Error(err error, msg ...string) {
	if err != nil {
		errMsg := err.Error()
		if len(msg) > 0 {
			errMsg = msg[0]
		}
		panic(errMsg)
	}
}
