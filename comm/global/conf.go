package global

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App  App  `json:"app" yaml:"app"`
	Keys Keys `json:"keys" yaml:"keys"`
}

type App struct {
	Env string `json:"env" yaml:"env"`
}

type Keys struct {
	BotName       string `json:"bot_name" yaml:"bot_name"`
	MasterAccount string `json:"master_account" yaml:"master_account"`
	MasterGroup   string `json:"master_group" yaml:"master_group"`
}

// GetConf .
func GetConf(cfg string) (conf *Config, err error) {
	var (
		yamlFile = make([]byte, 0)
	)

	filepath := fmt.Sprintf("%s", cfg)
	logrus.Infof("filepath: %s", filepath)
	yamlFile, err = os.ReadFile(filepath)
	if err != nil {
		err = errors.Wrapf(err, "ReadFile error")
		logrus.Errorf(err.Error())
		return conf, err
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		err = errors.Wrapf(err, "yaml.Unmarshal error")
		logrus.Errorf(err.Error())
		return conf, err
	}

	return conf, nil
}
