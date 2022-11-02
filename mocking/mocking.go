package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	executor func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.executor(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, "Go!")
}

func main() {
	Countdown(os.Stdout, &ConfigurableSleeper{1 * time.Second, time.Sleep})
}
