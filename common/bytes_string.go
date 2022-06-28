package common

import (
	"reflect"
	"unsafe"
)

//-------------------string与 []byte无拷贝转换------------------//
/*
 * 说明,在GO中 string存储结构为 reflect.StringHeader, Data 指针指向数据存储区域, Len表示数据长度
 * 而 slice 在存储结构为 reflect.SliceHeader, Data 指针指向数据存储区域, Len表示数据长度, Cap表示当前存储区的最大长度
 * slice与string存储上只相差一个Cap,且Data与Len的顺序是一致的。所以使用指针可以直接进行类型转换，而不需要内存拷贝,
 */



//convert string to  byte slice without copy
func StringBytes(s string) []byte {
	var bh reflect.SliceHeader
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// convert b to string without copy
func BytesString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

