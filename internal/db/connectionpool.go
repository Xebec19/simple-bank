package db

import (
	"context"
	"os"
	"sync"

	"github.com/Xebec19/simple-bank/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool
var once sync.Once
var query *Queries

func getDB(ctx context.Context) (*Queries, error) {

	var err error

	pool, err = pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	q := New(pool)

	return q, nil
}

func GetDB(ctx context.Context) (*Queries, error) {

	var err error

	once.Do(func() {
		query, err = getDB(ctx)
		if err != nil {
			logger.LogError("error: db connection failed ", err)
		}
	})

	return query, err
}

func ShutDownDB() {
	pool.Close()
}
