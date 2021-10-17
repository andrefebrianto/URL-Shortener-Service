package command

import (
	"context"
	"time"

	model "github.com/andrefebrianto/URL-Shortener-Service/model"
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
		PRIMARY_ID, shortlink.Code, shortlink.Url, shortlink.CreatedAt, shortlink.UpdatedAt, shortlink.ExpiredAt, shortlink.VisitorCounter).WithContext(ctx).Exec()
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

	err = session.Query("UPDATE shortlink set Url = ?, UpdatedAt = ?, ExpiredAt = ? WHERE Id = ? AND Code = ?;",
		shortlink.Url, time.Now().Local(), shortlink.ExpiredAt, PRIMARY_ID, shortlink.Code).WithContext(ctx).Exec()
	if err != nil {
		return err
	}

	return nil
}

func (repository CassandraCommandRepository) AddCounterByCode(ctx context.Context, code string, counter uint64) error {
	session, err := repository.cassandraClient.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	err = session.Query("UPDATE shortlink set visitorcounter = ? WHERE Id = ? AND Code = ?;",
		counter, PRIMARY_ID, code).WithContext(ctx).Exec()
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
