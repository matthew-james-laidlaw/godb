package assert

import (
	"reflect"
	"testing"
)

func ExpectEq(actual interface{}, expected interface{}, t *testing.T) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}