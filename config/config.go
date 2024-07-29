package config

import (
	"shortened-url/model"
	"sync"
)

var once sync.Once

type Cfg struct {
	Routes map[string]model.RouteRequest
}

var Config *Cfg

func Get() *Cfg {
	if Config == nil {
		once.Do(
			func() {
				Config = &Cfg{
					Routes: map[string]model.RouteRequest{},
				}
			})
	}

	return Config
}
