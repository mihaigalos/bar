package main

import "fmt"

type Bar struct {
	current  int64
	percent  int64
	progress string
	stop     int64
	symbol   string
}

func (bar *Bar) SetInitialProgress() {
	for i := 0; i < int(bar.percent); i += 2 {
		bar.progress += bar.symbol
	}
}

func (bar *Bar) New(start, stop int64) {
	bar.current = start
	bar.stop = stop
	bar.symbol = "█"
	bar.percent = bar.Percent()
	bar.SetInitialProgress()
}

func (bar *Bar) Percent() int64 {
	return int64(float32(bar.current) / float32(bar.stop) * 100)
}

func (bar *Bar) Update(current int64) {
	bar.current = current
	last := bar.percent
	bar.percent = bar.Percent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.progress += bar.symbol
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", bar.progress, bar.percent, bar.current, bar.stop)
}

func (bar *Bar) Finish() {
	fmt.Println()
}
