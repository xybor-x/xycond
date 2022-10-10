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

// Package xycond supports to assert or expect many conditions.
package xycond

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/xybor-x/xyerror"
)

type integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

type number interface {
	integer | float32 | float64
}

// failer instances may be *testing.T or *testing.B.
type failer interface {
	Fail()
}

type operator int

const (
	opEqual operator = iota
	opNotEqual
	opLessThan
	opNotLessThan
	opGreaterThan
	opNotGreaterThan
	opPanic
	opNil
	opNotNil
	opEmpty
	opNotEmpty
	opIs
	opIsNot
	opSame
	opNotSame
	opWritable
	opNotWritable
	opReadable
	opNotReadable
	opError
	opErrorNot
	opIn
	opNotIn
	opTrue
	opFalse
)

// ExpectEqual returns a true Condition if the two values are equal.
func ExpectEqual(a, b any) Condition {
	return Condition{result: a == b, op: opEqual, params: []any{a, b}}
}

// ExpectNotEqual returns a true Condition if the two values are not equal.
func ExpectNotEqual(a, b any) Condition {
	return ExpectEqual(a, b).revert(opNotEqual)
}

// ExpectLessThan returns a true Condition if the first parameter is less than
// the second.
func ExpectLessThan[t number](a, b t) Condition {
	return Condition{result: a < b, op: opLessThan, params: []any{a, b}}
}

// ExpectNotLessThan returns a true Condition if the first parameter is not less
// than the second.
func ExpectNotLessThan[t number](a, b t) Condition {
	return ExpectLessThan(a, b).revert(opNotLessThan)
}

// ExpectGreaterThan returns a true Condition if the first parameter is greater
// than the second.
func ExpectGreaterThan[t number](a, b t) Condition {
	return Condition{result: a > b, op: opGreaterThan, params: []any{a, b}}
}

// ExpectNotGreaterThan returns a true Condition if the first parameter is not
// greater than the second.
func ExpectNotGreaterThan[t number](a, b t) Condition {
	return ExpectGreaterThan(a, b).revert(opNotGreaterThan)
}

// ExpectPanic returns a true Condition if it found a panic with a correct data
// after calling the function.
func ExpectPanic(r any, f func()) (c Condition) {
	defer func() {
		var data = recover()
		if target, ok := r.(error); ok {
			if err, ok := data.(error); ok {
				c.result = errors.Is(err, target)
			} else {
				c.result = false
			}
		} else {
			c.result = data == r
		}
		c.op = opPanic
		c.params = []any{r, data}
	}()

	f()
	return
}

// ExpectZero returns a true Condition if the parameter is zero.
func ExpectZero[T number](a T) Condition {
	var zero T
	return ExpectEqual(a, zero)
}

// ExpectNotZero returns a true Condition if the parameter is not zero.
func ExpectNotZero[T number](a T) Condition {
	var zero T = 0
	return ExpectNotEqual(a, zero)
}

// ExpectNil returns a true Condition if the parameter is nil.
func ExpectNil(a any) Condition {
	var cond = Condition{result: false, op: opNil, params: []any{a}}

	if a == nil {
		cond.result = true
	} else {
		var va = reflect.ValueOf(a)
		var expect = ExpectIs(a, reflect.Chan, reflect.Func, reflect.Interface,
			reflect.Map, reflect.Pointer, reflect.Slice)
		if expect.result && va.IsNil() {
			cond.result = true
		}
	}

	return cond
}

// ExpectNotNil returns a true Condition if the parameter is not nil.
func ExpectNotNil(a any) Condition {
	return ExpectNil(a).revert(opNotNil)
}

// ExpectEmpty returns a true Condition if the parameter is an empty string,
// slice, array, or channel.
func ExpectEmpty(a any) Condition {
	var va = reflect.ValueOf(a)
	return Condition{
		result: va.Len() == 0,
		op:     opEmpty, params: []any{a, va.Kind()},
	}
}

// ExpectNotEmpty returns a true Condition if the parameter is not an empty
// string, slice, array, or channel.
func ExpectNotEmpty(a any) Condition {
	return ExpectEmpty(a).revert(opNotEmpty)
}

// ExpectIs returns a true Condition if value belongs to one of passed kinds.
func ExpectIs(v any, kinds ...reflect.Kind) Condition {
	var kindV = reflect.TypeOf(v).Kind()
	var cond = Condition{result: false, op: opIs}
	for i := range kinds {
		if kindV == kinds[i] {
			cond.result = true
		}
	}
	cond.params = []any{kindV, kinds}
	return cond
}

// ExpectIsNot returns a true Condition if value doesn't belong to any passed
// kinds.
func ExpectIsNot(v any, kinds ...reflect.Kind) Condition {
	return ExpectIs(v, kinds...).revert(opIsNot)
}

// ExpectSame returns a true Condition if parameters are the same type.
func ExpectSame(v ...any) Condition {
	var t0 = reflect.TypeOf(v[0])
	var cond = Condition{result: true, op: opSame}
	var types = []string{fmt.Sprint(t0)}
	for i := 1; i < len(v); i++ {
		var ti = reflect.TypeOf(v[i])
		if t0 != ti {
			cond.result = false
		}
		types = append(types, fmt.Sprint(ti))
	}
	cond.params = []any{v}
	return cond
}

// ExpectNotSame returns a true Condition if there is at least one value whose
// type is different from the rest.
func ExpectNotSame(v ...any) Condition {
	return ExpectSame(v...).revert(opNotSame)
}

// ExpectWritable returns a true Condition if the channel is writable.
func ExpectWritable(c any) Condition {
	AssertIs(c, reflect.Chan)
	var dir = reflect.TypeOf(c).ChanDir()
	return Condition{
		result: dir == reflect.BothDir || dir == reflect.SendDir,
		op:     opWritable,
	}
}

// ExpectNotWritable returns a true Condition if the channel is not writable.
func ExpectNotWritable(c any) Condition {
	return ExpectWritable(c).revert(opNotWritable)
}

// ExpectReadable returns a true Condition if the channel is readable.
func ExpectReadable(c any) Condition {
	AssertIs(c, reflect.Chan)
	var dir = reflect.TypeOf(c).ChanDir()
	return Condition{
		result: dir == reflect.BothDir || dir == reflect.RecvDir,
		op:     opReadable,
	}
}

// ExpectNotReadable returns a true Condition if the channel is not readable.
func ExpectNotReadable(c any) Condition {
	return ExpectReadable(c).revert(opNotReadable)
}

// ExpectError returns a true Condition if err belongs to one of the passed
// targets.
func ExpectError(err error, targets ...error) Condition {
	var cond = Condition{result: false, op: opError}
	for i := range targets {
		if errors.Is(err, targets[i]) {
			cond.result = true
		}
	}
	cond.params = []any{err, targets}
	return cond
}

// ExpectErrorNot returns a true Condition if the err doesn't belong to any
// targets.
func ExpectErrorNot(err error, targets ...error) Condition {
	return ExpectError(err, targets...).revert(opErrorNot)
}

// ExpectIn returns a true Condition if the element is in the object. The object
// must be an array, slice, string, or map.
func ExpectIn(elem any, obj any) Condition {
	AssertIs(obj, reflect.Array, reflect.Slice, reflect.String, reflect.Map)

	var objV = reflect.ValueOf(obj)
	var elemV = reflect.ValueOf(elem)

	var cond = Condition{
		result: false,
		op:     opIn,
	}

	switch objV.Kind() {
	case reflect.Map:
		AssertEqual(objV.Type().Key(), elemV.Type())
		cond.result = objV.MapIndex(elemV) != reflect.Value{}
		cond.params = []any{elem, "map"}
	case reflect.Slice, reflect.Array:
		AssertEqual(objV.Type().Elem(), elemV.Type())
		for i := 0; i < objV.Len(); i++ {
			if elem == objV.Index(i).Interface() {
				cond.result = true
				break
			}
		}
		cond.params = []any{elem, "array"}
	case reflect.String:
		AssertIs(elem, reflect.String, reflect.Int32, reflect.Uint8)
		switch elemV.Kind() {
		case reflect.Int32:
			cond.result = strings.ContainsRune(obj.(string), elem.(rune))
			cond.params = []any{strconv.QuoteRune(elem.(rune)), obj}
		case reflect.String:
			cond.result = strings.Contains(obj.(string), elem.(string))
			cond.params = []any{strconv.Quote(elem.(string)), obj}
		}
	}

	return cond
}

// ExpectNotIn returns a true Condition if the element is not in the object. The
// object must be an array, slice, string, or map.
func ExpectNotIn(object any, element any) Condition {
	return ExpectIn(object, element).revert(opNotIn)
}

// ExpectTrue returns true if the the parameter is true.
func ExpectTrue(b bool) Condition {
	return Condition{result: b, op: opTrue}
}

// ExpectFalse returns a true Condition if the parameter is false.
func ExpectFalse(b bool) Condition {
	return ExpectTrue(b).revert(opFalse)
}

// Panicf panics with a formatted string.
func Panicf(msg string, a ...any) any {
	panic(xyerror.AssertionError.Newf(msg, a...))
}

// Panic panics with default formatted objects.
func Panic(a ...any) any {
	panic(xyerror.AssertionError.New(a...))
}

// JustPanic panics immediately.
func JustPanic() {
	Panic("")
}

// Condition supports to perform actions on expectation.
type Condition struct {
	result bool
	op     operator
	params []any
}

// Test will call Fail method if it is a false Condition. It is used while
// testing, with *testing.T or *testing.B.
func (c Condition) Test(f failer) {
	if c.result {
		return
	}
	var _, fn, ln, ok = runtime.Caller(1)
	if ok {
		fmt.Printf("%s:%d: ", fn, ln)
	}
	fmt.Println(c.generateMessage())
	f.Fail()
}

// Assert prints the message and panics if it is a false Condition.
func (c Condition) Assert(msg string) {
	if !c.result {
		if msg != "" {
			Panic(msg)
		} else {
			Panic(c.generateMessage())
		}
	}
}

// Assertf prints the format message and panics if it is a false Condition.
func (c Condition) Assertf(msg string, a ...any) {
	if !c.result {
		Panic(fmt.Sprintf(msg, a...))
	}
}

// True is performed when Condition is true.
func (c Condition) True(f func()) Condition {
	if c.result {
		f()
	}
	return c
}

// False is performed when Condition is false.
func (c Condition) False(f func()) Condition {
	if !c.result {
		f()
	}
	return c
}

// revert returns the reverse Condition.
func (c Condition) revert(op operator) Condition {
	return Condition{result: !c.result, op: op, params: c.params}
}

func (c Condition) generateMessage() string {
	switch c.op {
	case opEqual:
		return fmt.Sprintf("%v != %v", c.params[0], c.params[1])
	case opNotEqual:
		return fmt.Sprintf("got the same value (%v)", c.params[0])
	case opLessThan:
		return fmt.Sprintf("%v is not less than %v", c.params[0], c.params[1])
	case opNotLessThan:
		return fmt.Sprintf("%v is less than %v", c.params[0], c.params[1])
	case opGreaterThan:
		return fmt.Sprintf("%v is not greater than %v", c.params[0], c.params[1])
	case opNotGreaterThan:
		return fmt.Sprintf("%v is greater than %v", c.params[0], c.params[1])
	case opPanic:
		if c.params[0] == nil {
			return fmt.Sprintf("expect no panic, but got %v", c.params[1])
		}
		return fmt.Sprintf("expect a panic of %v, but got %v",
			c.params[0], c.params[1])
	case opNil:
		return fmt.Sprintf("expect a nil value, but got %v", c.params[0])
	case opNotNil:
		return "expect a not nil value, but got nil"
	case opEmpty:
		return fmt.Sprintf("expect a empty %s, but got %v",
			c.params[1], c.params[0])
	case opNotEmpty:
		return fmt.Sprintf("expect a not empty %s, but got empty", c.params[1])
	case opIs:
		return fmt.Sprintf("expect a value in %v, but got %v",
			c.params[1], c.params[0])
	case opIsNot:
		return fmt.Sprintf("expect a value not in %v, but got %v",
			c.params[1], c.params[0])
	case opSame:
		var values = c.params[0].([]any)
		var types []reflect.Type
		for i := range values {
			var diff = true
			var vt = reflect.TypeOf(values[i])
			for j := range types {
				if vt == types[j] {
					diff = false
				}
			}
			if diff {
				types = append(types, vt)
			}
		}
		return fmt.Sprintf("expect values to be the same type, but got %v",
			types)
	case opNotSame:
		var values = c.params[0].([]any)
		return fmt.Sprintf(
			"expect values to be not the same type, but got only %v",
			reflect.TypeOf(values[0]))
	case opWritable:
		return "expect a wrtiable channel, but it's not"
	case opNotWritable:
		return "expect not a wrtiable channel, but it is"
	case opReadable:
		return "expect a readable channel, but it's not"
	case opNotReadable:
		return "expect not a readable channel, but it is"
	case opError:
		return fmt.Sprintf("expect a error in %v, but got %v",
			c.params[1], c.params[0])
	case opErrorNot:
		return fmt.Sprintf("expect a error not in %v, but got %v",
			c.params[1], c.params[0])
	case opIn:
		return fmt.Sprintf("%v NOT IN %v", c.params[0], c.params[1])
	case opNotIn:
		return fmt.Sprintf("%v IN %v", c.params[0], c.params[1])
	case opTrue:
		return "expect true, but got false"
	case opFalse:
		return "expect false, but got true"
	}
	panic("no available operator")
}
