package assert

import (
	"reflect"
	"runtime/debug"
	"testing"
)

func ExpectEq(actual interface{}, expected interface{}, t *testing.T) {
	if !reflect.DeepEqual(actual, expected) {
		debug.PrintStack()
		t.Errorf("expected %#v, got %#v", expected, actual)
	}
}
