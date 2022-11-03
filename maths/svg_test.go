package maths

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"
)

func TestSVGWriterAtMidnight(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			SimpleTime(0, 0, 0),
			Line{150, 150, 150, 60},
		},
		{
			SimpleTime(0, 0, 30),
			Line{150, 150, 150, 240},
		},
	}

	for _, c := range cases {
		t.Run(GetTestName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func containsLine(l Line, ls []Line) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}
	return false
}
