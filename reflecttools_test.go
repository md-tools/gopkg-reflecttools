package reflecttools

import (
	"testing"
)

func TestInit(t *testing.T) {
	type NK struct {
		C string
	}
	type K struct {
		A string
		B NK
		M map[string]string
	}
	k := K{}
	Init(&k)
	if k.M == nil {
		t.Errorf("Map was not initialized")
	}
}
