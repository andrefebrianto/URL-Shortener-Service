package query

import (
	"context"
	"errors"
	"time"

	model "github.com/andrefebrianto/URL-Shortener-Service/model"
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

	shortlinks := make([]model.ShortLink, 0)
	scanner := session.Query("SELECT * FROM shortlink;").WithContext(ctx).Consistency(gocql.One).Iter().Scanner()
	for scanner.Next() {
		var Id string
		var Code string
		var Url string
		var CreatedAt time.Time
		var UpdatedAt time.Time
		var ExpiredAt time.Time
		var VisitorCounter uint64

		err = scanner.Scan(&Id, &Code, &CreatedAt, &ExpiredAt, &UpdatedAt, &Url, &VisitorCounter)
		if err != nil {
			return nil, err
		}

		shortlinks = append(shortlinks, model.ShortLink{Id: Id, Code: Code, Url: Url, CreatedAt: CreatedAt, UpdatedAt: UpdatedAt, ExpiredAt: ExpiredAt, VisitorCounter: VisitorCounter})
	}

	if len(shortlinks) == 0 {
		return nil, errors.New("not found")
	}

	return shortlinks, nil
}

func (repository CassandraQueryRepository) GetByCode(ctx context.Context, code string) (*model.ShortLink, error) {
	session, err := repository.cassandraClient.CreateSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var Id string
	var Code string
	var Url string
	var CreatedAt time.Time
	var UpdatedAt time.Time
	var ExpiredAt time.Time
	var VisitorCounter uint64

	err = session.Query("SELECT * FROM shortlink WHERE Id = ? AND Code = ?;", PRIMARY_ID, code).WithContext(ctx).Consistency(gocql.One).Scan(&Id, &Code,
		&CreatedAt, &ExpiredAt, &UpdatedAt, &Url, &VisitorCounter)
	if err != nil {
		return nil, err
	}

	shortlink := &model.ShortLink{Id: Id, Code: Code, Url: Url, CreatedAt: CreatedAt, UpdatedAt: UpdatedAt, ExpiredAt: ExpiredAt, VisitorCounter: VisitorCounter}

	return shortlink, nil
}
