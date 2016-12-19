package shellescape_test

import (
	"testing"

	"github.com/alessio/shellescape"
)

func assertEqual(t *testing.T, s, expected string) {
	if s != expected {
		t.Errorf("%q (expected: %q)", s, expected)
	}
}

func TestEmptyString(t *testing.T) {
	s := shellescape.Quote("")
	expected := "''"
	assertEqual(t, s, expected)
}

func TestDoubleQuotedString(t *testing.T) {
	s := shellescape.Quote(`"double quoted"`)
	expected := `'"double quoted"'`
	assertEqual(t, s, expected)
}

func TestSingleQuotedString(t *testing.T) {
	s := shellescape.Quote(`'single quoted'`)
	expected := `''"'"'single quoted'"'"''`
	assertEqual(t, s, expected)
}

func TestUnquotedString(t *testing.T) {
	s := shellescape.Quote(`no quotes`)
	expected := `'no quotes'`
	assertEqual(t, s, expected)
}

func TestSingleInvalid(t *testing.T) {
	s := shellescape.Quote(`;`)
	expected := `';'`
	assertEqual(t, s, expected)
}

func TestAllInvalid(t *testing.T) {
	s := shellescape.Quote(`;${}`)
	expected := `';${}'`
	assertEqual(t, s, expected)
}
