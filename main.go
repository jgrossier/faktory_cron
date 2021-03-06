package main

import (
	"flag"

	"github.com/robfig/cron"
)

var (
	version     = "master"
	debug       bool
	config_path string
	config      *Config
	scheduler   *cron.Cron
)

func init() {
	flag.StringVar(&config_path, "config", "./crontab.yaml", "path to the configuration file")
	flag.BoolVar(&debug, "debug", false, "enable debug logs")
	flag.Parse()
}

func main() {
	configureLogging()
	logger.Infof("Cron for Faktory (version %s)", version)
	logger.Debugf("Reading config from file: %v", config_path)

	config = NewConfig(config_path)
	err := config.Update()
	if err != nil {
		logger.Fatalf("Error in config: %v", err)
	}

	scheduler = cron.New()
	for _, t := range config.Jobs {
		logger.Infof("Will run %v job (%v) every %v", t.Name, t.Type, t.Schedule)
		t.AddToScheduler()
	}
	logger.Infof("Loaded %d scheduled tasks from %v", len(config.Jobs), config_path)

	scheduler.Start()
	defer scheduler.Stop()

	// run forever
	select {}
}
