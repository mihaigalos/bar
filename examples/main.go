package main

import (
	"time"

	"github.com/mihaigalos/bar"
)

func main() {
	var bar bar.Bar
	bar.New(0, 100)
	for i := 0; i <= 100; i++ {
		time.Sleep(100 * time.Millisecond)
		bar.Update(i)
	}
	bar.Finish()
}
