package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		name          string
		input         interface{}
		expectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				name string
				city string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			var got []string
			Walk(test.input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.expectedCalls) {
				t.Errorf("got %v, want %v", got, test.expectedCalls)
			}
		})
	}
}
