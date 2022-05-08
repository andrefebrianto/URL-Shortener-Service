package config

import (
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfiguration struct {
	Port                   string
	Environment            string
	CassandraConfiguration CassandraConfiguration
	BussinessConfiguration BussinessConfiguration
}

type CassandraConfiguration struct {
	Hosts    []string
	Keyspace string
}

type BussinessConfiguration struct {
}

var v = viper.New()

func init() {
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.SetEnvPrefix("USS")
	v.AutomaticEnv()
	v.SetDefault("server.port", ":3000")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AddConfigPath(".")
	v.AddConfigPath("../")

	if err := v.ReadInConfig(); err != nil {
		log.Errorf("Failed to read the configuration file: %s", err)
	}
}

func Load() *AppConfiguration {
	return &AppConfiguration{
		Port:        v.GetString("server.port"),
		Environment: v.GetString("server.environment"),
		CassandraConfiguration: CassandraConfiguration{
			Hosts:    v.GetStringSlice("cassandra.hosts"),
			Keyspace: v.GetString("cassandra.keyspace"),
		},
	}
}
