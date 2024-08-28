package main

import (
	"fmt"
	"io"
	"os"
	"time"
)


func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, "Go!")
}

type ConfigurableSleeper struct {
  duration time.Duration
  sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep()  {
  c.sleep(c.duration)
}
