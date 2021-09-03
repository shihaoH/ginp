package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	f := s[0:1:5]
	//fmt.Println(s)
	//app(s)
	fmt.Printf("slice: %v, cap: %v, len: %v \n", f, cap(f), len(f))
	//f = append(f, f...)
	//f = append(f, f...)
	//fmt.Printf("slice: %v, cap: %v, len: %v \n", f, cap(f), len(f))

}

func app(arr []int) {
	for _, a := range arr {
		fmt.Println(reflect.Kind(a) == reflect.Ptr)
	}
}
