package main

import (
	"github.com/konveyor/kofini/cmd"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.GetRootCmd().Execute(); err != nil {
		logrus.Fatalf("Error: %v", err)
	}
}
