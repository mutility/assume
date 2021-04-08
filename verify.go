// +build !noverify

/*
The assume package allows for conditionally adding panics for test builds that
have little or no impact on your "release" builds.

Specify the build tag 'noverify' to elide the panics.
*/
package assume

import "fmt"

// True panics if passed false, unless -tags noverify, then returns v.
func True(v bool) bool {
	if !v {
		panic("got false")
	}
	return v
}

// Truef panics with message if passed false, unless -tags noverify, then returns v.
func Truef(v bool, format string, a ...interface{}) bool {
	if !v {
		panic(fmt.Sprintf(format, a...))
	}
	return v
}

// False panics if passed true, unless -tags noverify, then returns v.
func False(v bool) bool {
	if v {
		panic("got true")
	}
	return v
}

// Falsef panics with message if passed true, unless -tags noverify, then returns v.
func Falsef(v bool, format string, a ...interface{}) bool {
	if v {
		panic(fmt.Sprintf(format, a...))
	}
	return v
}

// Ok panics if passed an error, unless -tags noverify, then returns err.
func Ok(err error) error {
	if err != nil {
		panic("got error: " + err.Error())
	}
	return err
}

// Ok panics with message if passed an error, unless -tags noverify, then returns err.
func Okf(err error, format string, a ...interface{}) error {
	if err != nil {
		panic(fmt.Sprintf(format, a...))
	}
	return err
}

// Err panics if passed a nil error, unless -tags noverify, then returns err.
func Err(err error) error {
	if err == nil {
		panic("got nil error")
	}
	return err
}

// Errf panics with message if passed a nil error, unless -tags noverify, then returns err.
func Errf(err error, format string, a ...interface{}) error {
	if err == nil {
		panic(fmt.Sprintf(format, a...))
	}
	return err
}

// Unreached panics if called, unless -tags noverify.
func Unreached() {
	panic("reached")
}

// Unreachedf panics with message if called, unless -tags noverify.
func Unreachedf(format string, a ...interface{}) {
	panic(fmt.Sprintf(format, a...))
}
