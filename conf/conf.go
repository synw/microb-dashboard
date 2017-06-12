package conf

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/synw/microb-dashboard/types"
	"github.com/synw/terr"
)

func GetConf(dev bool) (*types.Conf, *terr.Trace) {
	name := "normal"
	if dev {
		name = "dev"
	}
	return getConf(name)
}

func getConf(name string) (*types.Conf, *terr.Trace) {
	// set some defaults for conf
	if name == "dev" {
		viper.SetConfigName("dev_config")
	} else {
		viper.SetConfigName("config")
	}
	viper.AddConfigPath(".")
	viper.SetDefault("domain", "localhost")
	viper.SetDefault("addr", "localhost:8090")
	// get the actual conf
	err := viper.ReadInConfig()
	if err != nil {
		var conf *types.Conf
		switch err.(type) {
		case viper.ConfigParseError:
			trace := terr.New("conf.getConf", err)
			return conf, trace
		default:
			err := errors.New("Unable to locate config file")
			trace := terr.New("conf.getConf", err)
			return conf, trace
		}
	}
	domain := viper.GetString("domain")
	addr := viper.GetString("addr")
	conf := types.Conf{domain, addr}
	return &conf, nil
}
