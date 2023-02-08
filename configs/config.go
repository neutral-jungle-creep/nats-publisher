package configs

import (
	"flag"
	"github.com/spf13/viper"
)

type Config struct {
	ClusterID string
	ClientID  string
	URL       string
	Subj      string
}

func NewConfig() *Config {
	return &Config{
		URL:       viper.GetString("nats.URL"),
		ClientID:  viper.GetString("nats.client"),
		ClusterID: viper.GetString("nats.cluster"),
		Subj:      viper.GetString("nats.subj"),
	}
}

func init() {
	var configPath, configFile string

	flag.StringVar(&configPath, "path", "configs", "Path to config file")
	flag.StringVar(&configFile, "config", "config", "Name of config file")
	flag.StringVar(&configPath, "p", "configs", "Path to config file")
	flag.StringVar(&configFile, "c", "config", "Name of config file")
	flag.Parse()

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
