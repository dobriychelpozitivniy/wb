package repository

import (
	"context"
	"time"
	"wbl0/pkg/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Postgres interface {
	Add(*model.Message) error
	ReadAll() ([]*model.Message, error)
	GetLastTime() (*time.Time, error)
}

type Local interface {
	Add(msg *model.Message)
	Read(id string) (*model.Message, error)
	AddMany(msgs []*model.Message)
}

type Repository struct {
	Postgres
	Local
}

func NewRepository(msgs map[string]*model.Message) (*Repository, error) {
	db, err := connectToDB()
	if err != nil {
		return nil, err
	}

	return &Repository{
		Local:    NewLocalRepository(msgs),
		Postgres: NewPostgresRepository(db),
	}, nil
}

func connectToDB() (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	con, err := pgxpool.Connect(ctx, "postgres://postgres:qwer@localhost:5432/wbl0")
	if err != nil {
		return nil, err
	}

	return con, nil
}
