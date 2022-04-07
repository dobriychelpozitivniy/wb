package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"wbl0/pkg/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresRepository struct {
	DB *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{DB: db}
}

func (r *PostgresRepository) GetLastTime() (*time.Time, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	sql := "SELECT created_at FROM wbl0 WHERE created_at=(SELECT MAX(created_at) FROM wbl0)"

	row := r.DB.QueryRow(ctx, sql)

	var t time.Time
	err := row.Scan(&t)
	fmt.Println(t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *PostgresRepository) Add(msg *model.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	msgJson, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	sql := "INSERT INTO wbl0 VALUES ($1, $2)"
	_, err = r.DB.Exec(ctx, sql, msg.OrderUID, msgJson)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) ReadAll() ([]*model.Message, error) {
	var msgs []*model.Message

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	sql := "SELECT order_jsonb FROM wbl0"
	rows, err := r.DB.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var m model.Message
		err = rows.Scan(&m)
		if err != nil {
			return nil, err
		}

		msgs = append(msgs, &m)
	}

	// Any errors encountered by rows.Next or rows.Scan will be returned here
	if rows.Err() != nil {
		return nil, err
	}

	return msgs, nil
}
