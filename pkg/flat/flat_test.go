// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package flat

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestFlatInt(t *testing.T) {
	in := 1
	out := Flat(in)
	assert.Equal(t, []int{1}, out)
}

func TestFlatString(t *testing.T) {
	in := "a"
	out := Flat(in)
	assert.Equal(t, []string{"a"}, out)
}

func TestFlatIntSlice(t *testing.T) {
	in := []int{1, 2, 3}
	out := Flat(in)
	assert.Equal(t, []int{1, 2, 3}, out)
}

func TestFlatStringSlice(t *testing.T) {
	in := []string{"a", "b", "c"}
	out := Flat(in)
	assert.Equal(t, []string{"a", "b", "c"}, out)
}

func TestFlatStrings(t *testing.T) {
	in := []interface{}{
		[]interface{}{
			"a",
			"b",
			"c",
		},
		[]interface{}{},
		"d",
		"e",
		"f",
	}
	out := Flat(in)
	assert.Equal(t, []interface{}{"a", "b", "c", "d", "e", "f"}, out)
}

func TestFlatDepthStrings(t *testing.T) {
	in := []interface{}{
		[]interface{}{
			"a",
			[]interface{}{
				"b",
				[]interface{}{
					"c",
					"d",
					"e",
				},
			},
		},
	}
	out := FlatDepth(in, 0, 1)
	assert.Equal(t, []interface{}{"a", "b", []interface{}{"c", "d", "e"}}, out)
}

func TestFlatComplex(t *testing.T) {
	in := []interface{}{
		[]interface{}{
			"a",
			"b",
			"c",
		},
		[]int{1, 2, 3},
		[]interface{}{},
		"d",
		"e",
		"f",
	}
	out := Flat(in)
	assert.Equal(t, []interface{}{"a", "b", "c", 1, 2, 3, "d", "e", "f"}, out)
}
