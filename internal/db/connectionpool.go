package db

import (
	"context"
	"os"

	"github.com/Xebec19/simple-bank/internal/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func GetDB(ctx context.Context) {

	var err error

	pool, err = pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		// todo return err here
	}

	q := db.New(pool)
}
