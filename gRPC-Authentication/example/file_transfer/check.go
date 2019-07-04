package main

import (
	"context"
	"fmt"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/widgets/gauge"
	"io"
	"os"
	"time"
)

func main() {
	//fileTransfer()
	//stringCheck()
	pic()
}
// playType indicates how to play a gauge.
type playType int

const (
	playTypePercent playType = iota
	playTypeAbsolute
)

// playGauge continuously changes the displayed percent value on the gauge by the
// step once every delay. Exits when the context expires.
func playGauge(ctx context.Context, g *gauge.Gauge, step int, delay time.Duration, pt playType) {
	progress := 0
	mult := 1

	ticker := time.NewTicker(delay)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			switch pt {
			case playTypePercent:
				if err := g.Percent(progress); err != nil {
					panic(err)
				}
			case playTypeAbsolute:
				if err := g.Absolute(progress, 100); err != nil {
					panic(err)
				}
			}

			progress += step * mult
			if progress > 100 || 100-progress < step {
				progress = 100
			} else if progress < 0 || progress < step {
				progress = 0
			}

			if progress == 100 {
				mult = -1
			} else if progress == 0 {
				mult = 1
			}

		case <-ctx.Done():
			return
		}
	}
}

func pic()  {
	t, err := termbox.New()
	if err != nil {
		panic(err)
	}
	defer t.Close()
	ctx, _ := context.WithCancel(context.Background())
	slim, err := gauge.New(
		gauge.Height(1),
		gauge.Border(linestyle.Light),
		gauge.BorderTitle("Percentage progress"),
	)
	if err != nil {
		panic(err)
	}
	go playGauge(ctx, slim, 10, 500*time.Millisecond, playTypePercent)

}

func stringCheck(){
	n := "nandakumar"

	fmt.Println(n[1:])
}

func fileTransfer(){
	// open input file
	fi, err := os.Open("./input.dmg")
	if err != nil {
		panic(err)
	}

	stat, err := fi.Stat()
	if err != nil {
		return
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	// open output file
	fo, err := os.Create("./output.zip")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// make a buffer to keep chunks that are read
	buf := make([]byte, stat.Size())
	for {
		// read a chunk
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := fo.Write(buf[:n]); err != nil {
			panic(err)
		}
	}
}