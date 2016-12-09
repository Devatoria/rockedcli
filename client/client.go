package client

import (
	"sync"

	"github.com/Devatoria/rocked"
	"github.com/spf13/viper"
)

var instance *rocked.Rocked
var once sync.Once

func Get() *rocked.Rocked {
	once.Do(func() {
		var err error
		instance, err = rocked.NewRocked(viper.GetString("docker"))
		if err != nil {
			panic(err)
		}
	})

	return instance
}
