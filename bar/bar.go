package bar

import "fmt"

type Bar struct {
	current  int
	percent  int
	progress string
	stop     int
	symbol   string
}

func (bar *Bar) SetInitialProgress() {
	for i := 0; i < int(bar.percent); i += 2 {
		bar.progress += bar.symbol
	}
}

func (bar *Bar) New(start, stop int) {
	bar.current = start
	bar.stop = stop
	bar.symbol = "█"
	bar.percent = bar.GetPercent()
	bar.SetInitialProgress()
}

func (bar *Bar) GetPercent() int {
	return int(float32(bar.current) / float32(bar.stop) * 100)
}

func (bar *Bar) GetCurrent() int {
	return bar.current
}

func (bar *Bar) Update(current int) {
	bar.current = current
	last := bar.percent
	bar.percent = bar.GetPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.progress += bar.symbol
	}
	fmt.Printf("\r[%-50s]%3d%%", bar.progress, bar.percent)
}

func (bar *Bar) Finish() {
	fmt.Println()
}
