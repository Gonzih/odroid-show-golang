# ODRDOI-SHOW golang client

## Installation

```
go get -u github.com/Gonzih/odroid-show-golang
```


## Usage

```go
package main

import (
	"log"
	"github.com/Gonzih/odroid-show-golang"
)

func main() {
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
	odroid.WriteString("this is how you write data to your board")

	err = odr.Sync() // will actualyl send buffer contents to the board

	if err != nil {
		log.Fatal(err)
	}
}
```
