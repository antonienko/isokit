package isokit

import (
	"context"
	"reflect"
	"testing"
)

func TestNewRoute(t *testing.T) {
	handler := func(ctx context.Context) {}

	testCases := map[string]struct {
		path         string
		expectedVars []string
	}{
		"Without vars": {
			"/this/is/a/path",
			[]string(nil),
		},
		"Single var": {
			"/this/{firstVar}/is/a/path",
			[]string{"firstVar"},
		},
		"Multiple vars": {
			"/this/is/{firstVar}/a/{secondVar}/path",
			[]string{"firstVar", "secondVar"},
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			r := NewRoute(tc.path, handler)
			if !reflect.DeepEqual(r.varNames, tc.expectedVars) {
				t.Errorf("wrong vars: expected %#v, got %#v", tc.expectedVars, r.varNames)
			}
		})
	}
}
