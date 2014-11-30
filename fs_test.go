package main

import (
	"reflect"
	"testing"
)

func TestSplitPath(t *testing.T) {
	path := "/Red/round.Apple"
	parts := SplitPath(path)
	if g, e := parts, []string{"red", "round", "apple"}; !reflect.DeepEqual(g, e) {
		t.Errorf("SplitIncorrect: %q != %q", g, e)
	}
}
