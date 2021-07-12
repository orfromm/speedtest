package main

import (
	l "github.com/sirupsen/logrus"
	"os"
	"speedtest/cmd"
)

func main() {
	cmd.Execute()
	l.Info("Main - done")
	os.Exit(0)
}
