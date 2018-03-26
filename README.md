# ODRDOID-SHOW golang client

![demo](https://raw.githubusercontent.com/Gonzih/odroid-status-screen-go/master/demo.jpg)

[Sample project](https://github.com/Gonzih/odroid-status-screen-go)

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
	board, err := odroid.NewOdroidShowBoard("/dev/ttyUSB0")

	if err != nil {
		log.Fatal(err)
	}

	board.Clear()
	board.ColorReset()
	board.WriteString("hello from golang!")
	board.Ln()
	board.Fg(odroid.ColorBlack)
	board.Bg(odroid.ColorRed)
	board.WriteString("this is how you write data to your board")

	err = board.Sync() // will actualyl send buffer contents to the board

	if err != nil {
		log.Fatal(err)
	}
}
```
