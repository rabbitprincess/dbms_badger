package util

import (
	"testing"
)

func BenchmarkSnowflakeID(b *testing.B) {
	f := NewSnowflake(1)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = f.Next()
		}
	})
}
