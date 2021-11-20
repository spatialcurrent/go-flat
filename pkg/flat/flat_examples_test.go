// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package flat

import (
	"fmt"
)

func ExampleFlat() {
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
	fmt.Printf("%#v\n", out)
	// Output: []interface {}{"a", "b", "c", 1, 2, 3, "d", "e", "f"}
}

func ExampleFlatDepth() {
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
	maxDepth := 1
	out := FlatDepth(in, 0, maxDepth)
	fmt.Printf("%#v\n", out)
	// Output: []interface {}{"a", "b", []interface {}{"c", "d", "e"}}
}
