package command

import (
	"github.com/gocql/gocql"
)

type CassandraCommandRepository struct {
	cassandraClient *gocql.ClusterConfig
}

func CreateCassandraCommandRepository(cassandraClient *gocql.ClusterConfig) CassandraCommandRepository {
	return CassandraCommandRepository{cassandraClient: cassandraClient}
}
