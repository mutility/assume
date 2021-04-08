// +build noverify

package assume_test

import (
	"errors"
	"testing"

	"github.com/mutility/assume"
)

var err = errors.New("Bam")

func TestSilentPasses(t *testing.T) {
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
	assume.Unreached()
	assume.Unreachedf("won't see %s", "this")
}

func TestSilentFails(t *testing.T) {
	if assume.True(false) {
		t.Error("didn't return false")
	}
	if assume.Truef(false, "won't see %s", "this") {
		t.Error("didn't return false")
	}
	if !assume.False(true) {
		t.Error("didn't return true")
	}
	if !assume.Falsef(true, "won't see %s", "this") {
		t.Error("didn't return true")
	}
	if nil == assume.Ok(err) {
		t.Error("didn't return err")
	}
	if nil == assume.Okf(err, "won't see %s", "this") {
		t.Error("didn't return err")
	}
	if nil != assume.Err(nil) {
		t.Error("didn't return nil")
	}
	if nil != assume.Errf(nil, "won't see %s", "this") {
		t.Error("didn't return nil")
	}
}
