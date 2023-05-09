package main

import (
	"log"

	"github.com/tum-dev/calendar-proxy/internal"
)

func main() {
	app := &internal.App{}
	log.Println(app.Run())
}
