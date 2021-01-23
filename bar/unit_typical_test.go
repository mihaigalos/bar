package bar

import (
	"testing"
	"time"
)

func TestSendWorks_whenTypical(t *testing.T) {
	expected := 100

	var bar Bar
	bar.New(0, 500)
	for i := 0; i <= 500; i++ {
		time.Sleep(1 * time.Millisecond)
		bar.Update(i)
	}
	bar.Finish()
	actual := bar.GetPercent()

	if actual != expected {
		t.Errorf("No Match: %d != %d", actual, expected)
	}
}
