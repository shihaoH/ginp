package ginp

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"sync"
)

var responderList []Responder
var onceRespList sync.Once

func getResponderList() []Responder {
	onceRespList.Do(func() {
		responderList = []Responder{(StringResponder)(nil),
			(JsonResponder)(nil),
		}
	})
	return responderList
}

type Responder interface {
	RespondTo() gin.HandlerFunc
}

func Convert(handler interface{}) gin.HandlerFunc {
	hRef := reflect.ValueOf(handler)
	for _, r := range responderList {
		rRef := reflect.ValueOf(r).Elem()
		if rRef.Type().ConvertibleTo(hRef.Type()) {
			rRef.Set(hRef)
			return rRef.Interface().(Responder).RespondTo()
		}
	}
	return nil
}

type StringResponder func(*gin.Context) string

func (s StringResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(200, s(context))
	}
}

type Json interface{}
type JsonResponder func(*gin.Context) Json

func (j JsonResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, j(context))
	}
}
