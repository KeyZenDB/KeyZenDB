package internal

import "testing"

func BenchmarkDb_Set(b *testing.B) {
	db := NewDb()
	for i := 0; i < b.N; i++ {
		db.Set("key", "value")
	}
}

func BenchmarkDb_Get(b *testing.B) {
	db := NewDb()
	db.Set("key", "value")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = db.Get("key")
	}
}
