package xycond_test

import (
	"fmt"
	"testing"

	"github.com/xybor-x/xycond"
)

var keys []string
var str string
var m map[string]any

func init() {
	m = make(map[string]any)
	for i := 0; i < 100000; i++ {
		var key = fmt.Sprintf("random_key_%d", i)
		keys = append(keys, key)
		str += key
		m[key] = nil
	}
}

func BenchmarkExpectIn(b *testing.B) {
	b.Run("map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xycond.ExpectIn(keys[i%len(keys)], m)
		}
	})
	b.Run("array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xycond.ExpectIn(keys[i%len(keys)], keys)
		}
	})
	b.Run("string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xycond.ExpectIn(keys[i%len(keys)], str)
		}
	})
	b.Run("rune", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xycond.ExpectIn(keys[i%len(keys)][0], str)
		}
	})
}
