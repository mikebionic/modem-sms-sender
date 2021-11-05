package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type config struct {
	ListenAddress string `json:"listen_address"`
}


func run() error {
	var configFile string
	var conf *config
	
	configFile = "config.json"
	conf, err = ReadConfig(configFile)
	if err != nil {
		log.WithError(err).WithField("config-file", configFile).Error("error loading configuration")
		return err
	}

	a := App{}
	a.Initialize()
	a.Run(conf.ListenAddress)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
