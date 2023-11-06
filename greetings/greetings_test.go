package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "salomao"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("salomao")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("salomao") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloName_IfEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}

}
