// +build !noverify

package assume_test

import (
	"errors"
	"testing"

	"github.com/mutility/assume"
)

var err = errors.New("Bam")

func TestLoudPasses(t *testing.T) {
	if !assume.True(true) {
		t.Error("didn't return true")
	}
	if !assume.Truef(true, "won't see %s", "this") {
		t.Error("didn't return true")
	}
	if assume.False(false) {
		t.Error("didn't return false")
	}
	if assume.Falsef(false, "won't see %s", "this") {
		t.Error("didn't return false")
	}
	if nil != assume.Ok(nil) {
		t.Error("didn't return nil")
	}
	if nil != assume.Okf(nil, "won't see %s", "this") {
		t.Error("didn't return nil")
	}
	if nil == assume.Err(err) {
		t.Error("didn't return err")
	}
	if nil == assume.Errf(err, "won't see %s", "this") {
		t.Error("didn't return err")
	}
}

func TestLoudFails(t *testing.T) {
	willPanic := func(fn func()) {
		defer func() {
			if recover() == nil {
				t.Error("did not panic")
			}
		}()
		fn()
	}
	willPanic(func() { assume.True(false) })
	willPanic(func() { assume.Truef(false, "won't see %s", "this") })
	willPanic(func() { assume.False(true) })
	willPanic(func() { assume.Falsef(true, "won't see %s", "this") })
	willPanic(func() { _ = assume.Ok(err) })
	willPanic(func() { _ = assume.Okf(err, "won't see %s", "this") })
	willPanic(func() { _ = assume.Err(nil) })
	willPanic(func() { _ = assume.Errf(nil, "won't see %s", "this") })
	willPanic(func() { assume.Unreached() })
	willPanic(func() { assume.Unreachedf("won't see %s", "this") })
}
