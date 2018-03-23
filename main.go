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

type OdroidShowBoard struct {
	Port   *serial.Port
	Buffer *bytes.Buffer
}

func (odroid *OdroidShowBoard) Sync() error {
	payload := odroid.Buffer.Bytes()
	_, err := odroid.Port.Write(payload)
	odroid.Buffer.Reset()

	return err
}

func (odroid *OdroidShowBoard) Write(b []byte) {
	_, err := odroid.Buffer.Write(b)

	if err != nil {
		log.Fatalf("Error writing to a buffer: %s", err)
	}
}

func (odroid *OdroidShowBoard) WriteString(s string) {
	odroid.Write([]byte(s))
}

func (odroid *OdroidShowBoard) Clear() error {
	odroid.WriteString("\033c")
	return odroid.Sync()
}

func (odroid *OdroidShowBoard) Ln() {
	odroid.WriteString("\012\015")
}

func (odroid *OdroidShowBoard) Color(color int) {
	payload := fmt.Sprintf("\033[%dm", color)
	odroid.Write([]byte(payload))
}

func (odroid *OdroidShowBoard) ColorReset() {
	odroid.Fg(ColorWhite)
	odroid.Bg(ColorBlack)
}

func (odroid *OdroidShowBoard) Fg(color int) {
	if color < 10 {
		odroid.Color(30 + color)
	} else {
		log.Printf("Ignoring color core: %d", color)
	}
}

func (odroid *OdroidShowBoard) Bg(color int) {
	if color < 10 {
		odroid.Color(40 + color)
	} else {
		log.Printf("Ignoring color core: %d", color)
	}
}

func NewOdroidShowBoard(path string) (*OdroidShowBoard, error) {
	var odroid OdroidShowBoard
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
	odroid, err := NewOdroidShowBoard("/dev/ttyUSB0")

	if err != nil {
		log.Fatal(err)
	}

	odroid.Clear()
	odroid.ColorReset()
	odroid.WriteString("hello from golang!")
	odroid.Ln()
	odroid.WriteString("and second line!")

	odroid.Sync()
	time.Sleep(time.Second)

	odroid.Ln()
	odroid.WriteString("test")

	odroid.Sync()
	time.Sleep(time.Second)

	odroid.Ln()
	odroid.WriteString("READ STUFF")
	odroid.Ln()
	odroid.WriteString("MORE STUFF")
	odroid.Ln()
	odroid.WriteString("and more?!!!1111")

	odroid.Sync()
	time.Sleep(time.Second)
}
