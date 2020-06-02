package main

import (
	"log"
	"os"
	"time"

	"github.com/rjeczalik/notify"
)

func main() {
	c := make(chan notify.EventInfo, 1)

	if err := notify.Watch(".", c, notify.All); err != nil {
		log.Fatal(err)
	}
	defer notify.Stop(c)

	for {
		select {
		case ei := <-c:
			info, _ := os.Stat(ei.Path())
			log.Printf("NEWPR: Duration for %s(%s): %s", ei.Path(), ei.Event().String(), time.Since(info.ModTime()))
		}
	}
}
