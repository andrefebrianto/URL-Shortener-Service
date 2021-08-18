package command

import (
	"context"

	model "github.com/andrefebrianto/URL-Shortener-Service/src/model"
	"github.com/gocql/gocql"
)

type CassandraCommandRepository struct {
	cassandraClient *gocql.ClusterConfig
}

var PRIMARY_ID = "SEA"

func CreateCassandraCommandRepository(cassandraClient *gocql.ClusterConfig) CassandraCommandRepository {
	return CassandraCommandRepository{cassandraClient: cassandraClient}
}

func (repository CassandraCommandRepository) Create(ctx context.Context, shortlink *model.ShortLink) error {
	session, err := repository.cassandraClient.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	err = session.Query("INSERT INTO shortlink (Id, Code, Url, CreatedAt, UpdatedAt, ExpiredAt, VisitorCounter) VALUES (?, ?, ?, ?, ?, ?, ?) IF NOT EXISTS;",
		shortlink.Id, shortlink.Code, shortlink.Url, shortlink.CreatedAt, shortlink.UpdatedAt, shortlink.ExpiredAt, shortlink.VisitorCounter).WithContext(ctx).Exec()
	if err != nil {
		return err
	}

	return nil
}

func (repository CassandraCommandRepository) UpdateByCode(ctx context.Context, shortlink *model.ShortLink) error {
	session, err := repository.cassandraClient.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()
	err = session.Query("UPDATE shortlink set Url = ?, UpdatedAt = ?, ExpiredAt = ?, VisitorCounter = ? IF NOT EXISTS;",
		shortlink.Url, gocql.TimeUUID().Timestamp(), shortlink.ExpiredAt, shortlink.VisitorCounter).WithContext(ctx).Exec()
	if err != nil {
		return err
	}

	return nil
}

func (repository CassandraCommandRepository) DeleteByCode(ctx context.Context, code string) error {
	session, err := repository.cassandraClient.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()
	err = session.Query("DELETE FROM shortlink WHERE Id = ? AND Code = ?", PRIMARY_ID, code).WithContext(ctx).Exec()
	if err != nil {
		return err
	}

	return nil
}
