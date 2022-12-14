// Copyright (c) 2022 xybor-x
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package xycond_test

import (
	"reflect"
	"testing"

	"github.com/xybor-x/xycond"
	"github.com/xybor-x/xyerror"
)

func TestAssert(t *testing.T) {
	defer func() {
		var r = recover()
		if r == nil {
			t.Fail()
		}
	}()

	xycond.AssertTrue(false)
}

func TestAssertEqual(t *testing.T) {
	xycond.AssertEqual(1, 1)
	xycond.AssertNotEqual(1, 2)
}

func TestAssertLessThan(t *testing.T) {
	xycond.AssertLessThan(1, 2)
	xycond.AssertNotLessThan(1, 0)
}

func TestAssertGreaterThan(t *testing.T) {
	xycond.AssertGreaterThan(1, 0)
	xycond.AssertNotGreaterThan(1, 1)
}

func TestAssertPanic(t *testing.T) {
	xycond.AssertPanic("", func() { panic("") })
	xycond.AssertPanic(nil, func() {})
}

func TestAssertZero(t *testing.T) {
	xycond.AssertZero(0)
	xycond.AssertNotZero(1)
}

func TestAssertNil(t *testing.T) {
	var x *int
	xycond.AssertNil(x)
	xycond.AssertNil(nil)

	var a = make([]int, 0)
	xycond.AssertNotNil(a)
	xycond.AssertNotNil(new(int))

	var err error = xyerror.AssertionError.New("foo")
	xycond.AssertNotNil(err)
}

func TestAssertEmpty(t *testing.T) {
	xycond.AssertEmpty("")
	xycond.AssertEmpty([]int{})
	xycond.AssertEmpty([]int{1, 2, 3}[0:0])

	xycond.AssertNotEmpty("a")
	xycond.AssertNotEmpty([]int{1})
	xycond.AssertNotEmpty([1]int{1})
}

func TestAssertIs(t *testing.T) {
	xycond.AssertIs(1, reflect.Int)
	xycond.AssertIsNot(1, reflect.String)
}

func TestAssertSame(t *testing.T) {
	xycond.AssertSame(1, 2)
	xycond.AssertSame(1, 2, 3, 4, 5)
	xycond.AssertSame(make(chan int), make(chan int))

	xycond.AssertNotSame(1, "a")
	xycond.AssertNotSame(1, '3')
	xycond.AssertNotSame("a", 1)
	xycond.AssertNotSame(1, 2, 3, "a")
	xycond.AssertNotSame([]int{1}, [1]int{1})
}

func TestAssertWritable(t *testing.T) {
	var receive = make(<-chan int)
	var both = make(chan int)
	var send = make(chan<- int)

	xycond.AssertWritable(both)
	xycond.AssertWritable(send)
	xycond.AssertNotWritable(receive)
}

func TestAssertReadable(t *testing.T) {
	var send = make(chan<- int)
	var both = make(chan int)
	var receive = make(<-chan int)

	xycond.AssertReadable(both)
	xycond.AssertReadable(receive)
	xycond.AssertNotReadable(send)
}

func TestAssertError(t *testing.T) {
	var err = xyerror.ValueError.New("")
	xycond.AssertError(err, xyerror.ValueError)
	xycond.AssertErrorNot(err, xyerror.AssertionError)
}

func TestAssertInWithMap(t *testing.T) {
	var m = map[int]string{1: "foo", 2: "bar"}
	xycond.AssertIn(1, m)
	xycond.AssertIn(2, m)
	xycond.AssertNotIn(3, m)
}

func TestAssertInWithArray(t *testing.T) {
	var a = []string{"foo", "bar"}
	xycond.AssertIn("foo", a)
	xycond.AssertIn("bar", a)
	xycond.AssertNotIn("buzz", a)
}

func TestAssertInWithString(t *testing.T) {
	var s = "foo bar"
	xycond.AssertIn("foo", s)
	xycond.AssertIn('b', s)
	xycond.AssertNotIn("buzz", s)
}

func TestAssertTrue(t *testing.T) {
	xycond.AssertTrue(true)
	xycond.AssertFalse(false)
}
