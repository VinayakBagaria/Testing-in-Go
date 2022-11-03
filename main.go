package main

import (
	"os"
	"testing-go/maths"
	"time"
)

func main() {
	maths.SVGWriter(os.Stdout, time.Now())
}
