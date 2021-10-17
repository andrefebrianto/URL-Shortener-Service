package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/spf13/viper"
)

var cluster *gocql.ClusterConfig
var config = viper.New()

func init() {
	config.SetConfigFile(`configs/configs.json`)
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func SetupConnection() {
	hosts := config.GetStringSlice("cassandra.hosts")
	cluster = gocql.NewCluster(hosts...)
	cluster.Keyspace = config.GetString("cassandra.keyspace")
	// cluster.Authenticator = gocql.PasswordAuthenticator{
	// 	Username: config.GetString("cassandra.user"),
	// 	Password: config.GetString("cassandra.password"),
	// }
}

//GetConnection ...
func GetConnection() *gocql.ClusterConfig {
	return cluster
}
