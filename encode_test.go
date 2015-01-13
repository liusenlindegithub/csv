package csv

import (
	"reflect"
	"testing"
)

type P struct {
	First string
	Last  string
}

type X struct {
	First string
}

func TestMarshal_without_a_slice(t *testing.T) {
	_, err := Marshal(simple{})

	if err == nil {
		t.Error("Non slice produced no error")
	}
}

func TestEncodeFieldValue(t *testing.T) {
	var encTests = []struct {
		val      interface{}
		expected string
		tag      string
	}{
		// Strings
		{"ABC", "ABC", ""},
		{byte(123), "123", ""},

		// Numerics
		{int(1), "1", ""},
		{float32(3.2), "3.2", ""},
		{uint32(123), "123", ""},
		{complex64(1 + 2i), "(+1+2i)", ""},

		// Boolean
		{true, "Yes", `true:"Yes" false:"No"`},
		{false, "No", `true:"Yes" false:"No"`},

		// TODO Array
		// Interface with Marshaler
		{P{"Jay", "Zee"}, "Jay Zee", ""},

		// Struct without Marshaler will produce nothing
		{X{"Jay"}, "", ""},
	}

	for _, test := range encTests {
		fv := reflect.ValueOf(test.val)
		st := reflect.StructTag(test.tag)
		res := encodeFieldValue(fv, st)

		if res != test.expected {
			t.Errorf("%s does not match %s", res, test.expected)
		}
	}

}
