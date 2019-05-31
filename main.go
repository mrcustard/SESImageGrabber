package main

// reads a csv file with numeric values and downloads images from a given url


import (
	"github.com/mrcustard/SESImageGrabber/commands"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}