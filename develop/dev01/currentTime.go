package currentTime

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func CurrentTime() (time.Time, error) {
	date, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		_, err2 := fmt.Fprintf(os.Stderr, "Error: %v", err)
		if err2 != nil {
			fmt.Printf("Error occured: %v", err2)
		}
		os.Exit(400)
	}
	return date, nil
}
