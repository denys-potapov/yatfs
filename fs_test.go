package main

import (
	"reflect"
	"testing"
)

func TestOrderTags(t *testing.T) {
	tags := OrderTags([]string{"a", "b", "c"})
	if tags["a"] != 255 || tags["b"] != 127 || tags["c"] != 0 {
		t.Errorf("OrderIncorrect %v", tags)
	}
}

func TestSplitPath(t *testing.T) {
	path := "/Red/round.Apple"
	parts := SplitPath(path)
	if g, e := parts, []string{"red", "round", "apple"}; !reflect.DeepEqual(g, e) {
		t.Errorf("SplitIncorrect: %q != %q", g, e)
	}
}
