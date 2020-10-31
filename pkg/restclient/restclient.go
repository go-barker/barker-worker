package restclient

import (
	"github.com/corporateanon/barker-worker/pkg/config"
	"github.com/go-resty/resty/v2"
)

func New(config *config.Config) *resty.Client {
	return resty.New().SetHostURL(config.BarkerUrl)
}
