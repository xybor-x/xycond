[![Xybor founder](https://img.shields.io/badge/xybor-huykingsofm-red)](https://github.com/huykingsofm)
[![Go Reference](https://pkg.go.dev/badge/github.com/xybor-x/xycond.svg)](https://pkg.go.dev/github.com/xybor-x/xycond)
[![GitHub Repo stars](https://img.shields.io/github/stars/xybor-x/xycond?color=yellow)](https://github.com/xybor-x/xycond)
[![GitHub top language](https://img.shields.io/github/languages/top/xybor-x/xycond?color=lightblue)](https://go.dev/)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/xybor-x/xycond)](https://go.dev/blog/go1.18)
[![GitHub release (release name instead of tag name)](https://img.shields.io/github/v/release/xybor-x/xycond?include_prereleases)](https://github.com/xybor-x/xycond/releases/latest)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/a8c3269dd8654796a09a898406997e96)](https://www.codacy.com/gh/xybor-x/xycond/dashboard?utm_source=github.com&utm_medium=referral&utm_content=xyplatform/xyerror&utm_campaign=Badge_Grade)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/a8c3269dd8654796a09a898406997e96)](https://www.codacy.com/gh/xybor-x/xycond/dashboard?utm_source=github.com&utm_medium=referral&utm_content=xyplatform/xyerror&utm_campaign=Badge_Coverage)
[![Go Report](https://goreportcard.com/badge/github.com/xybor-x/xycond)](https://goreportcard.com/report/github.com/xybor-x/xycond)

# Introduction

Package xycond supports to assert or expect many conditions.

It makes source code to be shorter and more readable by using inline commands.

# Features

This package has the following features:

-   Assert a condition, panic in case condition is false.
-   Expect a condition to occur and perform actions on this expectation.

# Example

```golang
xycond.AssertFalse(1 == 2)

var x int
xycond.AssertZero(x)

// Test a condition with *testing.T or *testing.B.
var t = &testing.T{}
xycond.ExpectEmpty("").Test(t)

// Perform actions on an expectation.
xycond.ExpectEqual(1, 2).
	True(func() {
		fmt.Printf("1 == 2")
	}).
	False(func() {
		fmt.Printf("1 != 2")
	})

// Output:
// 1 != 2
```
