package cassandra

import (
	"github.com/andrefebrianto/URL-Shortener-Service/config"
	"github.com/gocql/gocql"
)

var cluster *gocql.ClusterConfig

func SetupConnection(config config.CassandraConfiguration) {
	hosts := config.Hosts
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = config.Keyspace
}

//GetConnection ...
func GetConnection() *gocql.ClusterConfig {
	return cluster
}
