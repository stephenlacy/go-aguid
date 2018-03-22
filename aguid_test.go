package aguid

import (
	"strings"
	"testing"
)

func TestEmptyString(t *testing.T) {
	str, err := Aguid("")
	if err != nil {
		t.Fail()
	}
	if str == "" {
		t.Error("Empty string should return UUID.V4")
	}
}

func TestValidString(t *testing.T) {
	str, err := Aguid("String Here")
	if err != nil {
		t.Fail()
	}
	if str == "" {
		t.Error("Valid string should return aguid")
	}
	if len(str) != 36 {
		t.Error("Should be 36 characters")
	}

	s := strings.Split(str, "-")
	if len(s[0]) != 8 {
		t.Error("Part 0 needs to be 8 characters")
	}
	if len(s[1]) != 4 {
		t.Error("Part 1 needs to be 4 characters")
	}
	if len(s[2]) != 4 {
		t.Error("Part 2 needs to be 4 characters")
	}
	if len(s[3]) != 4 {
		t.Error("Part 3 needs to be 4 characters")
	}
	if len(s[4]) != 12 {
		t.Error("Part 4 needs to be 12 characters")
	}
	if len(s[0]) != 8 {
		t.Error("Part 0 needs to be 8 characters")
	}
}

func TestDuplication(t *testing.T) {
	str, _ := Aguid("base string")
	str2, _ := Aguid("base string")
	if str != str2 {
		t.Error("Strings do not match")
	}
}

func TestRandomString(t *testing.T) {
	str, err := Aguid("")
	if err != nil {
		t.Fail()
	}
	if len(str) != 36 {
		t.Error("Should return a valid random UUID")
	}
}
