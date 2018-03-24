package odroid

import (
	"log"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	odroid, err := NewOdroidShowBoard("/dev/ttyUSB0")

	if err != nil {
		log.Fatal(err)
	}

	odroid.Clear()
	odroid.ColorReset()
	odroid.WriteString("hello from golang!")
	odroid.Ln()
	odroid.WriteString("and second line!")
	odroid.Ln()
	odroid.WriteString("test")
	odroid.Ln()
	odroid.WriteString("READ STUFF")
	odroid.Ln()
	odroid.WriteString("MORE STUFF")
	odroid.Ln()
	odroid.WriteString("and more?!!!1111")
	odroid.WriteString("now we can write everything!")

	odroid.Sync()
	time.Sleep(time.Second)
}
