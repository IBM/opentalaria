package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/ibm/opentalaria/config"
	"github.com/ibm/opentalaria/logger"
	"github.com/ibm/opentalaria/plugins"
	// We start a web server only in localdev mode, which should't expose any sensitive information.
	// If we add some web APIs one day, this functionality has to be reviewed.
	// _ "expvar"
)

func initLogger(config *config.Config) {
	// print the log level before setting the log level handler so we can see what is set in case warn or error are set.
	logLevel := config.LogLevel
	slog.Info("Setting log level to " + logLevel.String())

	// initialize logger with level handler based on LOG_LEVEL env variable.
	// The default log level is Warn, if no env is set or the value is invalid.
	//
	// JSON Handler might be better suited for a cloud environment. Set it with LOG_FORMAT=json env variable
	var handler slog.Handler
	if config.LogFormat == "json" {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		handler = logger.NewCustomHandler(os.Stdout, nil)
	}

	logger := slog.New(logger.NewLevelHandler(logLevel, handler))

	slog.SetDefault(logger)
}

func main() {
	ps, err := plugins.New()
	if err != nil {
		log.Printf("error init plugin: %v", err)
	}
	err = ps.Call()
	if err != nil {
		log.Printf("error invoking Call(): %+v\n", err)
	}

	return

	// confFile := flag.String("c", "server.properties", "Path to config file. Default is server.properties")
	// flag.Parse()

	// // global config object that will be passed to all downstream APIs and methods
	// conf, err := config.NewConfig(*confFile)
	// if err != nil {
	// 	slog.Error("Error initializing broker", "err", err)
	// 	os.Exit(1)
	// }

	// initLogger(conf)

	// if conf.OTProfile == config.Localdev {
	// 	slog.Info(fmt.Sprintf("starting in local dev mode, listening on port :%d", conf.DebugServerPort))
	// 	// start a web server if we are in local dev mode
	// 	go http.ListenAndServe(fmt.Sprintf(":%d", conf.DebugServerPort), nil)
	// }

	// server := NewServer(conf)
	// server.Run()
}
