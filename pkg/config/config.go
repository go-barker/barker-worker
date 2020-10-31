package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	BarkerUrl string `validate:"required"`
}

func NewConfig() (*Config, error) {
	c := &Config{}
	v := viper.New()
	v.AutomaticEnv()
	c.BarkerUrl = v.GetString("barker_url")
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
