package odroid

import (
	"log"
	"testing"
)

func TestSimple(t *testing.T) {
	odr, err := NewOdroidShowBoard("/dev/ttyUSB0")

	if err != nil {
		log.Fatal(err)
	}

	odr.Clear()
	odr.ColorReset()
	odr.WriteString("hello from golang!")
	odr.Ln()
	odr.Fg(ColorBlack)
	odr.Bg(ColorRed)
	odr.WriteString("testing line support")

	err = odr.Sync()

	if err != nil {
		log.Fatal(err)
	}
}
