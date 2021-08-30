package assert

import (
	"reflect"
	"runtime/debug"
	"testing"
)

func ExpectNoErr(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}

func ExpectEq(actual interface{}, expected interface{}, t *testing.T) {
	if !reflect.DeepEqual(actual, expected) {
		debug.PrintStack()
		t.Errorf("expected %#v, got %#v", expected, actual)
	}
}
