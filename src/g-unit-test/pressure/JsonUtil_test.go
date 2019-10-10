package pressure

import "testing"

func BenchmarkEncodePersonToJsonFile(b *testing.B) {
	b.Log("benchmark BenchmarkEncodePersonToJsonFile starting...")
	b.ReportAllocs()
	p := Person{
		Name:   "Tom",
		Age:    2,
		Rmb:    268000,
		Gender: "Male",
		Hobby:  []string{"eat fish", "play with tom"},
	}
	for i := 0; i <= b.N; i++ {
		EncodePersonToJsonFile(&p,"a.json")
	}
}

func BenchmarkDecodeJsonFileToPerson(b *testing.B) {
	b.Log("benchmark BenchmarkDecodeJsonFileToPerson starting...")
	b.ReportAllocs()
	for i := 0; i <= b.N; i++ {
		DecodeJsonFileToPerson("a.json")
	}
}
