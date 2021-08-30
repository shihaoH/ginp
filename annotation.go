package ginp

import "reflect"

type Annotation interface {
	SetTag(tag reflect.StructTag)
}

var AnnotationList []Annotation

func init() {
	AnnotationList = make([]Annotation, 0)
	AnnotationList = append(AnnotationList, new(Value))
}

func IsAnnotation(t reflect.Type) bool {
	for _, item := range AnnotationList {
		if reflect.TypeOf(item) == t {
			return true
		}
	}
	return false
}

type Value struct {
	tag reflect.StructTag
}

func (v *Value) SetTag(tag reflect.StructTag) {
	v.tag = tag
}