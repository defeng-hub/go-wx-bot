package conf

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// Conf .
type Conf struct {
	App  App  `json:"app" yaml:"app"`
	Keys Keys `json:"keys" yaml:"keys"`
}

// App .
type App struct {
	Env string `json:"env" yaml:"env"`
}

type Keys struct {
	BotName       string `json:"bot_name" yaml:"bot_name"`
	MasterAccount string `json:"master_account" yaml:"master_account"`
}

// GetConf .
func GetConf(cfg string) (conf *Conf, err error) {
	var (
		yamlFile = make([]byte, 0)
	)

	filepath := fmt.Sprintf("%s", cfg)
	logrus.Infof("filepath: %s", filepath)
	yamlFile, err = ioutil.ReadFile(filepath)
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
