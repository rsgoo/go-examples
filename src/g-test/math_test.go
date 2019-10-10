package main

import "testing"

//性能测试
func BenchmarkGetSum(b *testing.B) {
	b.Log("BenchmarkGetSum Starting...")
	//报告内存开销
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetSum(25)
	}
}

func BenchmarkGetRecursiveSum(b *testing.B) {
	b.Log("BenchmarkGetRecursiveSum Starting...")
	//报告内存开销
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetRecursiveSum(25)
	}
}

func BenchmarkGetFibonacci(b *testing.B) {
	b.Log("BenchmarkGetFibonacci Starting...")
	//报告内存开销
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetFibonacci(25)
	}
}

func BenchmarkGetFibonacciRecursive(b *testing.B) {
	b.Log("BenchmarkGetFibonacciRecursive Starting...")
	//报告内存开销
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetFibonacciRecursive(25)
	}
}

func BenchmarkGetFibonacciClosure(b *testing.B) {
	b.Log("BenchmarkGetFibonacciClosure Starting...")
	//报告内存开销
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetFibonacciClosure()
	}
}

//goos: darwin
//goarch: amd64
//pkg: g-test
//循环次数		   耗时（单位纳秒）       内存开销          内存分配块（对象类型）
//5000000	       350 ns/op	       0 B/op	       0 allocs/op
//--- BENCH: BenchmarkGetSum-4
//    math_test.go:27: BenchmarkGetSum Starting...
//    math_test.go:27: BenchmarkGetSum Starting...
//    math_test.go:27: BenchmarkGetSum Starting...
//    math_test.go:27: BenchmarkGetSum Starting...
//    math_test.go:27: BenchmarkGetSum Starting...
//500000	      2959 ns/op	       0 B/op	       0 allocs/op
//--- BENCH: BenchmarkGetRecursiveSum-4
//    math_test.go:36: BenchmarkGetRecursiveSum Starting...
//    math_test.go:36: BenchmarkGetRecursiveSum Starting...
//    math_test.go:36: BenchmarkGetRecursiveSum Starting...
//    math_test.go:36: BenchmarkGetRecursiveSum Starting...
//500000	      2463 ns/op	    8192 B/op	       1 allocs/op
//--- BENCH: BenchmarkGetFibonacci-4
//    math_test.go:45: BenchmarkGetFibonacci Starting...
//    math_test.go:45: BenchmarkGetFibonacci Starting...
//    math_test.go:45: BenchmarkGetFibonacci Starting...
//    math_test.go:45: BenchmarkGetFibonacci Starting...
