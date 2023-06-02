package database

import (
	"context"
	"fmt"
	"pdi-go-kafka-bd/internal/utils"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DbConnection interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Close()
}

type DbPool struct {
	conn *pgxpool.Pool
}

func GetConnection(ctx context.Context) *DbPool {
	config, err := utils.LoadConfig()
	if err != nil {
		fmt.Println(err)
	} else {
		db_url := strings.Trim(fmt.Sprint(config), "{}")
		pool, err := pgxpool.Connect(ctx, db_url)
		if err != nil {
			fmt.Println("Unable to connect to database:", err)

			return nil
		}
		return &DbPool{
			conn: pool,
		}
	}
	return nil
}

func (pg *DbPool) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return pg.conn.Query(ctx, sql, args...)
}

func (pg *DbPool) Close() {
	pg.conn.Close()
}
