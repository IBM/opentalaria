package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/ibm/opentalaria/api"
	"github.com/ibm/opentalaria/config"
	"github.com/ibm/opentalaria/kafka"
	"github.com/ibm/opentalaria/protocol"

	// We start a web server only in localdev mode, which should't expose any sensitive information.
	// If we add some web APIs one day, this functionality has to be reviewed.
	_ "expvar"
)

func main() {
	confFile := flag.String("c", "config.yaml", "Path to config file. Default is config.yaml")
	flag.Parse()

	// global config object that will be passed to all downstream APIs and methods
	conf, err := config.NewConfig(*confFile)
	if err != nil {
		slog.Error("Error initializing broker", "err", err)
		os.Exit(1)
	}

	if conf.OTProfile == config.Localdev {
		slog.Info(fmt.Sprintf("starting in local dev mode, listening on port :%d", conf.DebugServerPort))
		// start a web server if we are in local dev mode
		go http.ListenAndServe(fmt.Sprintf(":%d", conf.DebugServerPort), nil)
	}

	server := kafka.NewServer(conf)

	server.RegisterAPI(&protocol.ApiVersionsRequest{}, 0, 3, api.HandleAPIVersionsRequest)
	server.RegisterAPI(&protocol.MetadataRequest{}, 0, 8, api.HandleMetadataRequest)
	server.RegisterAPI(&protocol.CreateTopicsRequest{}, 0, 4, api.HandleCreateTopics)
	server.RegisterAPI(&protocol.DeleteTopicsRequest{}, 0, 4, api.HandleDeleteTopics)

	server.Run()
}
