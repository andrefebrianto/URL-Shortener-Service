package config

import (
	"fmt"
	"strings"

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
	v.SetConfigType("yml")
	v.SetEnvPrefix("USS")
	v.AutomaticEnv()
	v.SetDefault("port", ":3000")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AddConfigPath(".")
	v.AddConfigPath("../")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Failed to read the configuration file: %s", err)
	}
}

func Load() *AppConfiguration {
	return &AppConfiguration{
		Port:        v.GetString("port"),
		Environment: v.GetString("environment"),
		CassandraConfiguration: CassandraConfiguration{
			Hosts:    v.GetStringSlice("cassandra-hosts"),
			Keyspace: v.GetString("cassandra-keyspace"),
		},
	}
}
