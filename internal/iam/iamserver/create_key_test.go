package iamserver

import (
	"regexp"
	"testing"
)

func TestGetKeyId(t *testing.T) {
	keyid := getKeyId(10)

	r, _ := regexp.MatchString("[A-z0-9]{10}", keyid)

	if !r {
		t.Errorf("Error getkeyid function not working")
	}
}
