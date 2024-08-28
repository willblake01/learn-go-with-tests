package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// ConfigurableSleeper is an implementation of Sleeper with a defined delay.
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleeper allows you to put delays.
type Sleeper interface {
	Sleep()
}

// Sleep will pause execution for the defined Duration.
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const finalWord = "Go!"
const countdownStart = 3

// Countdown prints a countdown from 3 to Go! with a 1 second pause between each number.
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintf(out, "%d\n", i)
		sleeper.Sleep()
	}
	fmt.Fprintf(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
