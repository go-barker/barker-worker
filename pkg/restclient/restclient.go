package restclient

import (
	"log"

	"github.com/corporateanon/barker-worker/pkg/config"
	"github.com/go-resty/resty/v2"
)

func New(config *config.Config) *resty.Client {
	log.Printf("Connecting to Barker instance via %s\n", config.BarkerUrl)
	return resty.New().SetHostURL(config.BarkerUrl)
}
