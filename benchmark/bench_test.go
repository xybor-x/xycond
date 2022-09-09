package benchmark

import (
	"fmt"
	"testing"

	"github.com/xybor-x/xycond"
)

var largeKeys []string
var largeString string
var largeMap map[string]any

var smallKeys []string
var smallString string
var smallMap map[string]any

func init() {
	largeMap = make(map[string]any)
	smallMap = make(map[string]any)
	for i := 0; i < 100000; i++ {
		var key = fmt.Sprintf("random_key_%d", i)
		largeKeys = append(largeKeys, key)
		largeString += key
		largeMap[key] = nil
	}
	for i := 0; i < 9; i++ {
		var key = fmt.Sprintf("random_key_%d", i)
		smallKeys = append(smallKeys, key)
		smallString += key
		smallMap[key] = nil
	}
}

func BenchmarkExpectIn(b *testing.B) {
	b.Run("large-map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xycond.ExpectIn(largeKeys[i%len(largeKeys)], largeMap)
		}
	})
	b.Run("small-map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xycond.ExpectIn(smallKeys[i%len(smallKeys)], smallMap)
		}
	})
	b.Run("large-array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xycond.ExpectIn(largeKeys[i%len(largeKeys)], largeKeys)
		}
	})
	b.Run("small-array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xycond.ExpectIn(smallKeys[i%len(smallKeys)], smallKeys)
		}
	})
	b.Run("large-string-string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xycond.ExpectIn(largeKeys[i%len(largeKeys)], largeString)
		}
	})
	b.Run("small-string-string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xycond.ExpectIn(smallKeys[i%len(smallKeys)], smallString)
		}
	})
	b.Run("large-string-rune", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xycond.ExpectIn(largeKeys[i%len(largeKeys)][0], largeString)
		}
	})
	b.Run("small-string-rune", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xycond.ExpectIn(smallKeys[i%len(smallKeys)][0], smallString)
		}
	})
}
