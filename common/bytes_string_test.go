package common

import (
	"testing"
)

func Test_StringBytes(t *testing.T) {
	s := "hello world!"
	b := StringBytes(s)
	t.Log(b)
	t.Logf("len:%d cap:%d", len(b), cap(b))

	bb := []byte(s)
	t.Log(bb)

	t.Logf("len:%d cap:%d", len(bb), cap(bb))
}

func Test_BytesString(t *testing.T) {
	b := []byte("hello world!")

	s := BytesString(b)
	t.Log(s)

	ss := string(b)
	t.Log(ss)
}

func Benchmark_BytesString(b *testing.B) {
	var x = []byte("hello world!")
	for i := 0; i < b.N; i++ {
		_ = BytesString(x)
	}
}

func Benchmark_Normal(b *testing.B) {
	var x = []byte("hello world!")
	for i := 0; i < b.N; i++ {
		_ = string(x)
	}
}

//go test -bench="."
//Benchmark_BytesString-4         1000000000               0.338 ns/op
//Benchmark_Normal-4              212061979                5.13 ns/op
//从性能测试看,使用with copy版本的函数比常规版本的函数速度快了1个数量级以上
