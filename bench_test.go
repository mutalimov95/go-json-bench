package jsonbench

import (
	"testing"

	simpeJSON "github.com/bitly/go-simplejson"
	insaneJSON "github.com/vitkovskii/insane-json"
)

func BenchmarkUnmarshal(b *testing.B) {
	b.Run("simple-json", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = simpeJSON.NewJson(tenFieldsByte)
		}
	})
	b.Run("insane-json", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			j, _ := insaneJSON.DecodeBytes(tenFieldsByte)
			insaneJSON.Release(j)
		}
	})
}
