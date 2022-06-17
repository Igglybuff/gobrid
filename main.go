package main

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/Igglybuff/gobrid/pkg/realdebrid"

	"fyne.io/fyne/v2/app"
)

func main() {
	ctx := context.Background()
	log.Info("Starting up...")
	a := app.New()
	version := "v0.0.1"
	title := fmt.Sprintf("gobrid %s", version)
	w := a.NewWindow(title)

	realdebrid.Load(ctx, w)

	w.Show()

	a.Run()

	log.Info("Tidying up...")

	tidyUp()
}

func tidyUp() {
	log.Info("Exited")
}
