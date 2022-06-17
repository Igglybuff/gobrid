package main

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Window struct {
	fyne.Window
}

func main() {
	ctx := context.Background()
	log.Info("Starting up...")
	a := app.New()
	version := "v0.0.1"
	title := fmt.Sprintf("gobrid %s", version)
	w := Window{a.NewWindow(title)}

	// load real-debrid integration
	RdLoad(ctx, &w)

	w.Show()

	a.Run()

	log.Info("Tidying up...")

	tidyUp()
}

func tidyUp() {
	log.Info("Exited")
}
