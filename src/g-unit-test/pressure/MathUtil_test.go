package pressure

import "testing"

func BenchmarkMyGetSum(b *testing.B) {
	b.Log("BenchmarkMyGetSum Starting...")
	//报告内存开销
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		PressureGetSum(25)
	}
}

func BenchmarkPressureGetRecursiveSum(b *testing.B) {
	b.Log("BenchmarkPressureGetRecursiveSum  Starting...")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		PressureGetRecursiveSum(25)
	}
}
