package poker

import (
	"os"
)

type Tape struct {
	File *os.File
}

func (t *Tape) Write(p []byte) (int, error) {
	t.File.Truncate(0)
	t.File.Seek(0, 0)
	return t.File.Write(p)
}
