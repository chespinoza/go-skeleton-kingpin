package main

import (
	"log"
	"os"

	"github.com/chespinoza/go-skeleton-kingpin/server"
	"go.uber.org/zap"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	// General Config
	kingpin.Version("0.0.1")
	startCmd := kingpin.Command("start", "start server")

	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	logger, err := zap.NewProduction(
		zap.Fields(
			zap.String("host", hostName),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Parse Commands
	switch kingpin.Parse() {
	case "start":
		config := server.Config{
			Parameter1: *startCmd.Flag("param1", "param1 config").OverrideDefaultFromEnvar("PARAM1").Required().String(),
			Parameter2: *startCmd.Flag("param2", "param2 config").Default("false").Bool(),
			Parameter3: *startCmd.Flag("param3", "param3 config").Default("1").Int64(),
		}
		s := server.New(&config, logger)
		s.Run()

	case "ping":
		logger.Info("Pong")
	}
}
