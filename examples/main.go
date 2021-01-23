package main

import (
	"time"

	"github.com/mihaigalos/go-bar/bar"
)

func main() {
	var bar bar.Bar
	bar.New(0, 248)
	for i := 0; i <= 248; i++ {
		time.Sleep(10 * time.Millisecond)
		bar.Update(i)
	}
	bar.Finish()
}
