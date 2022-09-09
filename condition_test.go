package xycond_test

import (
	"reflect"
	"testing"

	"github.com/xybor-x/xycond"
	"github.com/xybor-x/xyerror"
)

type mocktest struct{}

func (mocktest) Fail() {}

func TestCondition(t *testing.T) {
	xycond.ExpectTrue(false).Test(mocktest{})
	xycond.ExpectTrue(false).True(func() {
		t.Fail()
	})
	xycond.ExpectTrue(true).False(func() {
		t.Fail()
	})
}

func TestExpectEqual(t *testing.T) {
	xycond.ExpectEqual(1, 1).Test(t)
	xycond.ExpectNotEqual(1, 2).Test(t)
}

func TestExpectLessThan(t *testing.T) {
	xycond.ExpectLessThan(1, 2).Test(t)
	xycond.ExpectNotLessThan(1, 0).Test(t)
	xycond.ExpectNotLessThan(1, 1).Test(t)
}

func TestExpectGreaterThan(t *testing.T) {
	xycond.ExpectGreaterThan(1, 0).Test(t)
	xycond.ExpectNotGreaterThan(1, 1).Test(t)
	xycond.ExpectNotGreaterThan(1, 2).Test(t)
}

func TestExpectPanic(t *testing.T) {
	xycond.ExpectPanic("", func() { panic("") }).Test(t)
	xycond.ExpectPanic(nil, func() {}).Test(t)
}

func TestExpectZero(t *testing.T) {
	xycond.ExpectZero(0).Test(t)
	xycond.ExpectNotZero(1).Test(t)
}

func TestExpectNil(t *testing.T) {
	var x *int
	xycond.ExpectNil(x).Test(t)
	xycond.ExpectNil(nil).Test(t)

	var a = make([]int, 0)
	xycond.ExpectNotNil(a).Test(t)
	xycond.ExpectNotNil(new(int)).Test(t)

	var err error = xyerror.AssertionError.New("foo")
	xycond.ExpectNotNil(err).Test(t)
}

func TestExpectEmpty(t *testing.T) {
	xycond.ExpectEmpty("").Test(t)
	xycond.ExpectEmpty([]int{}).Test(t)
	xycond.ExpectEmpty([]int{1, 2, 3}[0:0]).Test(t)

	xycond.ExpectNotEmpty("a").Test(t)
	xycond.ExpectNotEmpty([]int{1}).Test(t)
	xycond.ExpectNotEmpty([1]int{1}).Test(t)
}

func TestExpectIs(t *testing.T) {
	var tests = map[any]reflect.Kind{
		1:     reflect.Int,
		"foo": reflect.String,
		1.1:   reflect.Float64,
		true:  reflect.Bool,
		'c':   reflect.Int32,
	}

	for value, kind := range tests {
		xycond.ExpectIs(value, kind).Test(t)
		xycond.ExpectIsNot(value, kind+1).Test(t)
	}
}

func TestExpectSame(t *testing.T) {
	xycond.ExpectSame(1, 2).Test(t)
	xycond.ExpectSame(1, 2, 3, 4, 5).Test(t)
	xycond.ExpectSame(make(chan int), make(chan int)).Test(t)

	xycond.ExpectNotSame(1, "a").Test(t)
	xycond.ExpectNotSame(1, '3').Test(t)
	xycond.ExpectNotSame("a", 1).Test(t)
	xycond.ExpectNotSame(1, 2, 3, "a").Test(t)
	xycond.ExpectNotSame([]int{1}, [1]int{1}).Test(t)
}

func TestExpectWritable(t *testing.T) {
	var receive = make(<-chan int)
	var both = make(chan int)
	var send = make(chan<- int)

	xycond.ExpectWritable(both).Test(t)
	xycond.ExpectWritable(send).Test(t)
	xycond.ExpectNotWritable(receive).Test(t)
}

func TestExpectReadable(t *testing.T) {
	var send = make(chan<- int)
	var both = make(chan int)
	var receive = make(<-chan int)

	xycond.ExpectReadable(both).Test(t)
	xycond.ExpectReadable(receive).Test(t)
	xycond.ExpectNotReadable(send).Test(t)
}

func TestExpectError(t *testing.T) {
	var err = xyerror.ValueError.New("")
	xycond.ExpectError(err, xyerror.ValueError).Test(t)
	xycond.ExpectErrorNot(err, xyerror.AssertionError).Test(t)
}

func TestExpectInWithMap(t *testing.T) {
	var m = map[int]string{1: "foo", 2: "bar"}
	xycond.ExpectIn(1, m).Test(t)
	xycond.ExpectIn(2, m).Test(t)
	xycond.ExpectNotIn(3, m).Test(t)
}

func TestExpectInWithArray(t *testing.T) {
	var a = []string{"foo", "bar"}
	xycond.ExpectIn("foo", a).Test(t)
	xycond.ExpectIn("bar", a).Test(t)
	xycond.ExpectNotIn("buzz", a).Test(t)
}

func TestExpectInWithString(t *testing.T) {
	var s = "foo bar"
	xycond.ExpectIn("foo", s).Test(t)
	xycond.ExpectIn('b', s).Test(t)
	xycond.ExpectNotIn("buzz", s).Test(t)
}

func TestExpectInInvalid(t *testing.T) {
	xycond.ExpectPanic(xyerror.AssertionError, func() {
		xycond.ExpectIn(1, 2)
	}).Test(t)
	xycond.ExpectPanic(xyerror.AssertionError, func() {
		xycond.ExpectIn(1, []string{})
	}).Test(t)
	xycond.ExpectPanic(xyerror.AssertionError, func() {
		xycond.ExpectIn(1, map[float32]int{})
	}).Test(t)
	xycond.ExpectPanic(xyerror.AssertionError, func() {
		xycond.ExpectIn(1, "")
	}).Test(t)
}

func TestExpectTrue(t *testing.T) {
	xycond.ExpectTrue(true).Test(t)
	xycond.ExpectFalse(false).Test(t)
}

func TestExpectAssert(t *testing.T) {
	xycond.ExpectPanic(xyerror.AssertionError, func() {
		xycond.ExpectFalse(true).Assert("")
	}).Test(t)
	xycond.ExpectPanic(xyerror.AssertionError, func() {
		xycond.ExpectFalse(true).Assertf("")
	}).Test(t)
}

func TestPanic(t *testing.T) {
	xycond.ExpectPanic(xyerror.AssertionError, func() {
		xycond.Panic("")
	}).Test(t)
	xycond.ExpectPanic(xyerror.AssertionError, func() {
		xycond.Panicf("")
	}).Test(t)
	xycond.ExpectPanic(xyerror.AssertionError, func() {
		xycond.JustPanic()
	}).Test(t)
}
