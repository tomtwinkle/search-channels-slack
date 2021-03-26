package config

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

const configFile = "slacktools.yaml"

type ConfigYaml struct {
	Slack *ConfigSlack `yaml:"slack"`
}

type ConfigSlack struct {
	Token string `yaml:"token"`
}

type config struct {
	ConfigPath string
}

type Config interface {
	Init() (*ConfigYaml, error)
	Read() (*ConfigYaml, error)
}

func NewConfig() Config {
	return &config{ConfigPath: getConfigPath()}
}

func getConfigPath() string {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	execDirPath := filepath.Dir(execPath)
	return filepath.Join(execDirPath, configFile)
}

func (c *config) Init() (*ConfigYaml, error) {
	if _, err := os.Stat(c.ConfigPath); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		return c.writeConfig()
	}
	return c.readConfig()
}

func (c *config) Read() (*ConfigYaml, error) {
	if _, err := os.Stat(c.ConfigPath); err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("config not found. please execute init command\n$ slacktool init")
		} else {
			return nil, err
		}
	}
	return c.readConfig()
}

func (c *config) readConfig() (*ConfigYaml, error) {
	file, err := os.Open(c.ConfigPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	var cfg ConfigYaml
	if err := d.Decode(&cfg); err != nil {
		return nil, err
	}

	if cfg.Slack.Token == "" {
		return nil, fmt.Errorf("slack token is Required [%s]", c.ConfigPath)
	}
	return &cfg, nil
}

func (c *config) writeConfig() (*ConfigYaml, error) {
	token, err := c.inputSlackToken()
	if err != nil {
		return nil, err
	}
	configSlack := &ConfigSlack{
		Token: token,
	}

	file, err := os.Create(c.ConfigPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	e := yaml.NewEncoder(file)
	defer e.Close()
	cfg := ConfigYaml{
		Slack: configSlack,
	}
	if err := e.Encode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
