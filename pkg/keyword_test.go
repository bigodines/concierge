package keyword

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	res := New("foobar, baz", "", true)
	if len(res.excludes) != 0 {
		fmt.Printf("%s", len(res.excludes))
		t.Error("Exclude words should be empty")
	}

	if "foobar" != res.keywords[0] || "baz" != res.keywords[1] {
		t.Error("Didn't parse keywords properly")
	}

	// one more happy path with case insensitive
	res = New("SOME, THING", "notME, norME", false)
	if len(res.excludes) != 2 || len(res.keywords) != 2 {
		t.Error("Didn't parse input properly")
	}

	if "some" != res.keywords[0] {
		t.Error("Should have converted words to lower case because case sensitive is false")
	}
}

func TestCheckAll(t *testing.T) {
	var res bool
	matcher := New("game, tonight", "was", false)

	res = matcher.CheckAll("wanna do go the game tonight")
	if !res {
		t.Error("Failed to match keywords from input")
	}
}
