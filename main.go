package main

import (
	"time"
)

func main() {
	var bar Bar
	bar.New(0, 100)
	for i := 0; i <= 100; i++ {
		time.Sleep(100 * time.Millisecond)
		bar.Update(int64(i))
	}
	bar.Finish()
}
