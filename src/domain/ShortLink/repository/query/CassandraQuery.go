package query

import (
	"github.com/gocql/gocql"
)

type CassandraQueryRepository struct {
	cassandraClient *gocql.ClusterConfig
}

func CreateCassandraQueryRepository(cassandraClient *gocql.ClusterConfig) CassandraQueryRepository {
	return CassandraQueryRepository{cassandraClient: cassandraClient}
}
