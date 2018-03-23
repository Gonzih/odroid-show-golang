package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

const (
	ColorBlack   = 0
	ColorRed     = 1
	ColorGreen   = 2
	ColorYellow  = 3
	ColorBlue    = 4
	ColorMagenta = 5
	ColorCyan    = 6
	ColorWhite   = 7
	ColorDefault = 9
)

type OdroidShow struct {
	Port   *serial.Port
	Buffer *bytes.Buffer
}

func (odroid *OdroidShow) Flush() error {
	payload := odroid.Buffer.Bytes()
	_, err := odroid.Port.Write(payload)
	odroid.Buffer.Reset()

	return err
}

func (odroid *OdroidShow) Write(b []byte) {
	_, err := odroid.Buffer.Write(b)

	if err != nil {
		log.Fatalf("Error writing to a buffer: %s", err)
	}
}

func (odroid *OdroidShow) WriteString(s string) {
	odroid.Write([]byte(s))
}

func (odroid *OdroidShow) Clear() error {
	_, err := odroid.Port.Write([]byte("\033c"))
	return err
}

func (odroid *OdroidShow) Ln() {
	odroid.WriteString("\012\015")
}

func (odroid *OdroidShow) Color(color int) {
	payload := fmt.Sprintf("\033[%dm", color)
	odroid.Write([]byte(payload))
}

func (odroid *OdroidShow) ColorReset() {
	odroid.Fg(ColorWhite)
	odroid.Bg(ColorBlack)
}

func (odroid *OdroidShow) Fg(color int) {
	odroid.Color(30 + color)
}

func (odroid *OdroidShow) Bg(color int) {
	odroid.Color(40 + color)
}

func InitOdroidShow(path string) (*OdroidShow, error) {
	var odroid OdroidShow
	var buffer bytes.Buffer

	conf := &serial.Config{Name: path, Baud: 500000}
	serialPort, err := serial.OpenPort(conf)

	if err != nil {
		return &odroid, err
	}

	odroid.Buffer = &buffer
	odroid.Port = serialPort
	return &odroid, nil
}

func main() {
	odroid, err := InitOdroidShow("/dev/ttyUSB0")

	if err != nil {
		log.Fatal(err)

	}

	odroid.Clear()
	odroid.ColorReset()
	odroid.WriteString("hello from golang!")
	odroid.Ln()
	odroid.WriteString("and second line!")

	odroid.Flush()
	time.Sleep(time.Second)

	odroid.Ln()
	odroid.WriteString("test")

	odroid.Flush()
	time.Sleep(time.Second)

	odroid.Ln()
	odroid.WriteString("READ STUFF")
	odroid.Ln()
	odroid.WriteString("MORE STUFF")
	odroid.Ln()
	odroid.WriteString("and more?!!!1111")

	odroid.Flush()
	time.Sleep(time.Second)
}
