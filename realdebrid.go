package main

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

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

func RdLoad(ctx context.Context, w *Window) {
	rdClient := DebridConnection{RealDebridStatus: "Disconnected"}

	rdTitle := widget.NewLabel("Real-Debrid integration")
	rdStatusLabel := widget.NewLabel("Checking...")

	rdConnectButton := widget.NewButton("Connect", func() {
		status := rdClient.RealDebridStatus
		if status == "Connected" {
			rdClient.RealDebridClient = nil
			rdClient.RealDebridStatus = "Disconnected"
		} else {
			_, err := rdClient.rdConnect(ctx)
			if err != nil {
				rdStatusLabel.SetText(rdClient.RealDebridStatus)
			} else {
				rdStatusLabel.SetText(err.Error())
			}
		}
	})

	rdIntegrationContent := container.NewVBox(
		rdTitle,
		rdStatusLabel,
		rdConnectButton,
	)

	go func() {
		for {
			time.Sleep(time.Second)
			updateButtonText(&rdClient, rdConnectButton)
			if rdStatusLabel.Text == "Checking..." || rdClient.RealDebridStatus == "Disconnected" {
				rdStatusLabel.SetText("Not connected")
			}
			w.SetContent(rdIntegrationContent)
		}
	}()
}

func (c *DebridConnection) rdConnect(ctx context.Context) (*realdebrid.User, error) {
	auth := realdebrid.Auth{KeyOrToken: rdToken}
	c.RealDebridClient = realdebrid.NewClient(realdebrid.DefaultClientOpts, auth, nil)
	user, err := c.RealDebridClient.GetUser(ctx)
	if err != nil {
		c.RealDebridStatus = "Connected"
	} else {
		msg := fmt.Sprintf("Failed to connect to RealDebrid!\n %s", err)
		log.Errorf(msg)
		c.RealDebridStatus = "Disconnected"
	}

	return &user, err
}

func updateButtonText(c *DebridConnection, button *widget.Button) {
	if c.RealDebridStatus == "Connected" {
		button.SetText("Disconnect")
	} else {
		button.SetText("Connect")
	}
	button.Refresh()
}
