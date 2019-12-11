package config

import "github.com/jinzhu/configor"

type Config struct {
	AppName string `default:"apptest"`
	Port    string `default:"8080"`
	DB      struct {
		Use      string `default:"postgres"`
		Postgres struct {
			Enabled  bool   `default:"true"`
			Host     string `default:"localhost"`
			Port     string `default:"5432"`
			UserName string `default:"postgres"`
			Password string `default:"postgres"`
			Database string `default:"apptest-db"`
		}
	}
	Contacts struct {
		Name  string `default:"Andres Matos"`
		Email string `default:"andres.matos@spsa.pe"`
	}
}

func NewConfig() (*Config, error) {
	c := &Config{}
	err := configor.Load(c, "config.yml")
	if err != nil {
		return nil, err
	}
	return c, nil
}
