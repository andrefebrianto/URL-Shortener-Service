package query

import (
	"context"
	"fmt"

	model "github.com/andrefebrianto/URL-Shortener-Service/src/model"
	"github.com/gocql/gocql"
)

type CassandraQueryRepository struct {
	cassandraClient *gocql.ClusterConfig
}

var PRIMARY_ID = "SEA"

func CreateCassandraQueryRepository(cassandraClient *gocql.ClusterConfig) CassandraQueryRepository {
	return CassandraQueryRepository{cassandraClient: cassandraClient}
}

func (repository CassandraQueryRepository) GetAll(ctx context.Context) ([]model.ShortLink, error) {
	session, err := repository.cassandraClient.CreateSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	err = session.Query("SELECT * FROM shortlink;").WithContext(ctx).Consistency(gocql.One).Exec()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return nil, nil
}

func (repository CassandraQueryRepository) GetByCode(ctx context.Context, code string) (*model.ShortLink, error) {
	session, err := repository.cassandraClient.CreateSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	shortlink := &model.ShortLink{}

	err = session.Query("SELECT * FROM shortlink WHERE Id = ? AND Code = ?;", PRIMARY_ID, code).WithContext(ctx).Consistency(gocql.One).Scan(shortlink)
	if err != nil {
		return nil, err
	}

	return shortlink, nil
}
