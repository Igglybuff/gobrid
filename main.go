package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/deflix-tv/go-debrid/realdebrid"
)

const (
	rdToken = ""
)

type DebridConnection struct {
	RealDebridClient *realdebrid.Client
	RealDebridStatus string
}

func main() {
	log.Info("Starting up...")
	a := app.New()
	version := "v0.0.1"
	title := fmt.Sprintf("gobrid %s", version)
	w := a.NewWindow(title)

	rdClient := DebridConnection{}
	rdClient.RealDebridStatus = "Disconnected"

	rdStatusLabel := widget.NewLabel(rdClient.RealDebridStatus)
	w.SetContent(container.NewVBox(
		rdStatusLabel,
		widget.NewButton("Connect", func() {
			err := rdClient.rdConnect()
			if err != nil {
				rdStatusLabel.SetText("Failed to connect to RealDebrid!")
			} else {
				rdStatusLabel.SetText(rdClient.RealDebridStatus)
			}
		}),
	))

	w.Show()
	a.Run()
	tidyUp()
}

func tidyUp() {
	log.Info("Exited")
}

func (c *DebridConnection) rdConnect() error {
	auth := realdebrid.Auth{KeyOrToken: rdToken}
	rd := realdebrid.NewClient(realdebrid.DefaultClientOpts, auth, nil)
	c.RealDebridClient = rd
	c.RealDebridStatus = "Connected!"
	return nil
}
