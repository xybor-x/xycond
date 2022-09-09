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

	xycond.ExpectTrue(false).
		True(t.Fail).
		False(func() {})
	xycond.ExpectTrue(true).
		True(func() {}).
		False(t.Fail)
}

func TestXxx(t *testing.T) {
	var x *int
	var tests = []xycond.Condition{
		xycond.ExpectTrue(false),
		xycond.ExpectFalse(true),
		xycond.ExpectEqual(1, 2),
		xycond.ExpectNotEqual(1, 1),
		xycond.ExpectLessThan(1, 0),
		xycond.ExpectNotLessThan(1, 2),
		xycond.ExpectGreaterThan(1, 2),
		xycond.ExpectNotGreaterThan(1, 0),
		xycond.ExpectPanic("", func() {}),
		xycond.ExpectPanic(nil, func() { panic("") }),
		xycond.ExpectZero(1),
		xycond.ExpectNotZero(0),
		xycond.ExpectNil(new(int)),
		xycond.ExpectNil(make([]int, 0)),
		xycond.ExpectNil(xyerror.AssertionError.New("foo")),
		xycond.ExpectNotNil(nil),
		xycond.ExpectNotNil(x),
		xycond.ExpectEmpty("a"),
		xycond.ExpectEmpty([]int{1}),
		xycond.ExpectEmpty([1]int{1}),
		xycond.ExpectNotEmpty(""),
		xycond.ExpectNotEmpty([]int{}),
		xycond.ExpectNotEmpty([]int{1, 2, 3}[0:0]),
		xycond.ExpectIs(3, reflect.String),
		xycond.ExpectIsNot(3, reflect.Int),
		xycond.ExpectSame(1, "a"),
		xycond.ExpectSame(1, '3'),
		xycond.ExpectSame("a", 1),
		xycond.ExpectSame("a", 1),
		xycond.ExpectSame(1, 2, 3, "a"),
		xycond.ExpectSame([]int{1}, [1]int{1}),
		xycond.ExpectNotSame(1, 2),
		xycond.ExpectNotSame(1, 2, 3, 4, 5),
		xycond.ExpectNotSame(make(chan int), make(chan int)),
		xycond.ExpectWritable(make(<-chan int)),
		xycond.ExpectNotWritable(make(chan int)),
		xycond.ExpectNotWritable(make(chan<- int)),
		xycond.ExpectReadable(make(chan<- int)),
		xycond.ExpectNotReadable(make(chan int)),
		xycond.ExpectNotReadable(make(<-chan int)),
		xycond.ExpectError(xyerror.ValueError.New(""), xyerror.KeyError),
		xycond.ExpectErrorNot(xyerror.ValueError.New(""), xyerror.ValueError),
		xycond.ExpectIn(3, map[int]string{1: "foo"}),
		xycond.ExpectIn("buzz", []string{"foo"}),
		xycond.ExpectIn("buzz", "foo bar"),
		xycond.ExpectNotIn(1, map[int]string{1: "foo"}),
		xycond.ExpectNotIn("foo", []string{"foo"}),
		xycond.ExpectNotIn("foo", "foo bar"),
	}

	for i := range tests {
		xycond.ExpectPanic(xyerror.AssertionError, func() {
			tests[i].Assert("")
		}).Test(t)
	}
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
		xycond.ExpectFalse(true).Assert("foo")
	}).Test(t)
	xycond.ExpectPanic(xyerror.AssertionError, func() {
		xycond.ExpectFalse(true).Assertf("foo")
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
