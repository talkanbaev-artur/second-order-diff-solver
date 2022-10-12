package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

type config struct {
	Port int
}

var cfg config = config{Port: 8000}

func Config() config {
	return cfg
}

func Init() {
	defer func() {
		if err := recover(); err != nil {
			logrus.Fatalf("Could not initiate the configuration, err: %v", err)
		}
	}()
	port, ok := readIntOpt("PORT")
	if ok {
		cfg.Port = port
	}
}

func readOpt(name string) (string, bool) {
	env, ok := os.LookupEnv(name)
	return env, ok
}

func readIntOpt(name string) (int, bool) {
	env, ok := os.LookupEnv(name)
	if !ok {
		return 0, false
	}
	val, err := strconv.Atoi(env)
	if err != nil {
		panic(fmt.Errorf("Failed to conver string env val to int, err: %s", err.Error()))
	}
	return val, ok
}
