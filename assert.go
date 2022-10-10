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

package xycond

import "reflect"

// AssertEqual panics if a is different from b.
func AssertEqual(a, b any) {
	ExpectEqual(a, b).Assert("")
}

// AssertNotEqual panics if a is equal to b.
func AssertNotEqual(a, b any) {
	ExpectNotEqual(a, b).Assert("")
}

// AssertLessThan panics if a is not less than b.
func AssertLessThan[t number](a, b t) {
	ExpectLessThan(a, b).Assert("")
}

// AssertNotLessThan panics if a is less than b.
func AssertNotLessThan[t number](a, b t) {
	ExpectNotLessThan(a, b).Assert("")
}

// AssertGreaterThan panics if a is not greater than b.
func AssertGreaterThan[t number](a, b t) {
	ExpectGreaterThan(a, b).Assert("")
}

// AssertNotGreaterThan panics if a is greater than b.
func AssertNotGreaterThan[t number](a, b t) {
	ExpectNotGreaterThan(a, b).Assert("")
}

// AssertPanic panics if the function doesn't panic.
func AssertPanic(r any, f func()) {
	ExpectPanic(r, f).Assert("")
}

// AssertZero panics if the parameter is not zero.
func AssertZero[t number](a t) {
	ExpectZero(a).Assert("")
}

// AssertNotZero panics if the parameter is zero.
func AssertNotZero[t number](a t) {
	ExpectNotZero(a).Assert("")
}

// AssertNil panics if the parameter is not nil.
func AssertNil(a any) {
	ExpectNil(a).Assert("")
}

// AssertNotNil panics if the parameter is nil.
func AssertNotNil(a any) {
	ExpectNotNil(a).Assert("")
}

// AssertEmpty panics if the parameter is not empty.
func AssertEmpty(a any) {
	ExpectEmpty(a).Assert("")
}

// AssertNotEmpty panics if the parameter is empty.
func AssertNotEmpty(a any) {
	ExpectNotEmpty(a).Assert("")
}

// AssertIs panics if value doesn't belongs to any passed kinds.
func AssertIs(v any, kinds ...reflect.Kind) {
	ExpectIs(v, kinds...).Assert("")
}

// AssertIsNot panics if value belongs to one of passed kinds.
func AssertIsNot(v any, kinds ...reflect.Kind) {
	ExpectIsNot(v, kinds...).Assert("")
}

// AssertSame panics if there is at least value' type different from the rest.
func AssertSame(v ...any) {
	ExpectSame(v...).Assert("")
}

// AssertNotSame panics if all values' type are the same.
func AssertNotSame(v ...any) {
	ExpectNotSame(v...).Assert("")
}

// AssertWritable panics if the parameter is not a writable channel.
func AssertWritable(c any) {
	ExpectWritable(c).Assert("")
}

// AssertNotWritable panics if the parameter is a writable channel.
func AssertNotWritable(c any) {
	ExpectNotWritable(c).Assert("")
}

// AssertReadable panics if the parameter is not a readable channel.
func AssertReadable(c any) {
	ExpectReadable(c).Assert("")
}

// AssertNotReadable panics if the parameter is a readable channel.
func AssertNotReadable(c any) {
	ExpectNotReadable(c).Assert("")
}

// AssertError panics if the err doesn't belong to any targets.
func AssertError(err error, targets ...error) {
	ExpectError(err, targets...).Assert("")
}

// AssertErrorNot panics if the err belongs to one of targets.
func AssertErrorNot(err error, targets ...error) {
	ExpectErrorNot(err, targets...).Assert("")
}

// AssertIn panics if the element is in the object. The object must be an array,
// slice, string, or map.
func AssertIn(element any, object any) {
	ExpectIn(element, object).Assert("")
}

// AssertNotIn panics if the element is not in the object. The object must be an
// array, slice, string, or map.
func AssertNotIn(element any, object any) {
	ExpectNotIn(element, object).Assert("")
}

// AssertTrue panics if the condition is false.
func AssertTrue(b bool) {
	ExpectTrue(b).Assert("")
}

// AssertFalse panics if the condition is true.
func AssertFalse(b bool) {
	ExpectFalse(b).Assert("")
}
