package currentTime

import (
	"fmt"
	"testing"
	"time"
)

func TestCurrentTime(t *testing.T) {
	for i := 0; i < 5; i++ {
		expected := time.Now()
		actual, err := CurrentTime()
		if err != nil {
			t.Error(err)
		}

		delta := actual.Sub(expected)
		if delta > time.Second {
			t.Errorf("Got wrong time: %v", delta)
		}
		fmt.Println("delta: ", delta)
	}
}
