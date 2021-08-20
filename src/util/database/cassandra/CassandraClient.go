package cassandra

import (
	"fmt"

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
	hosts := config.GetStringSlice("cassandraHosts")
	fmt.Println(hosts)
	cluster = gocql.NewCluster(hosts...)
	// cluster.Authenticator = gocql.PasswordAuthenticator{
	// 	Username: "user",
	// 	Password: "password",
	// }
}

//GetConnection ...
func GetConnection() *gocql.ClusterConfig {
	fmt.Println(cluster)
	return cluster
}
