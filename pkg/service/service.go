package service

import (
	"time"
	"wbl0/pkg/model"
	"wbl0/pkg/repository"
)

type Messages interface {
	AddMessage(*model.Message) error
	RestoreCash() error
	GetMessage(id string) (*model.Message, error)
	GetLastTime() (*time.Time, error)
}

type Service struct {
	Messages
}

func NewService(r *repository.Repository) *Service {
	return &Service{Messages: NewMessagesService(r.Postgres, r.Local)}
}
