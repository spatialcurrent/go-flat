// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// Package flat provides functions to recursively flatten a slice of slices.
package flat

import (
	"reflect"
	"time"
)

// Flat recursively flattens a slice of slices to a maximum depth.
// Flat attempts to maintain the slice type if possible.
// If maxDepth is less than zero, then there is no maximum depth.
func FlatDepth(in interface{}, depth int, maxDepth int) interface{} {

	// If a literal, then return a slice of that literal.
	switch in := in.(type) {
	case string:
		return []string{in}
	case int:
		return []int{in}
	case int64:
		return []int64{in}
	case float64:
		return []float64{in}
	case byte:
		return []byte{in}
	case time.Time:
		return []time.Time{in}
	}

	// If a literal slice or array, then return the slice directly
	switch in.(type) {
	case []string, []int, []int64, []float64, []byte, []time.Time:
		return in
	}

	// If reached max depth.
	if maxDepth >= 0 && depth > maxDepth {
		return in
	}

	v := reflect.ValueOf(in)
	t := v.Type()
	k := t.Kind()

	if k == reflect.Array || k == reflect.Slice {
		if v.Len() == 0 {
			return in
		}
		if v.Len() == 1 {
			return FlatDepth(v.Index(0).Interface(), depth+1, maxDepth)
		}
		out := reflect.MakeSlice(t, 0, v.Len())
		for i := 0; i < v.Len(); i++ {
			vi := v.Index(i).Interface()
			viv := reflect.ValueOf(vi)
			ti := viv.Type()
			if ki := ti.Kind(); ki == reflect.Array || ki == reflect.Slice {
				x := reflect.ValueOf(FlatDepth(vi, depth+1, maxDepth))
				if x.Type().AssignableTo(t) {
					out = reflect.AppendSlice(out, x)
				} else {
					for j := 0; j < x.Len(); j++ {
						out = reflect.Append(out, x.Index(j))
					}
				}
				continue
			}
			out = reflect.Append(out, viv)
		}
		return out.Interface()
	}

	return reflect.Append(reflect.MakeSlice(reflect.SliceOf(t), 0, 1), v).Interface()
}

// Flat recursively flattens a slice of slices.
// Flat attempts to maintain the slice type if possible.
func Flat(in interface{}) interface{} {
	return FlatDepth(in, 0, -1)
}
