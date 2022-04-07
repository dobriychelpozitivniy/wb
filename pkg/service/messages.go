package service

import (
	"time"
	"wbl0/pkg/model"
	"wbl0/pkg/repository"

	"github.com/jackc/pgx/v4"
)

type MessagesService struct {
	pgrsRepo repository.Postgres
	cashRepo repository.Local
}

func NewMessagesService(pgrs repository.Postgres, cash repository.Local) *MessagesService {
	return &MessagesService{pgrsRepo: pgrs, cashRepo: cash}
}

func (s *MessagesService) AddMessage(msg *model.Message) error {
	err := s.pgrsRepo.Add(msg)
	if err != nil {
		return err
	}

	s.cashRepo.Add(msg)

	return nil
}

func (s *MessagesService) GetLastTime() (*time.Time, error) {
	return s.pgrsRepo.GetLastTime()
}

func (s *MessagesService) RestoreCash() error {
	msgs, err := s.pgrsRepo.ReadAll()
	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	if err != pgx.ErrNoRows {
		s.cashRepo.AddMany(msgs)
	}

	return nil
}

func (s *MessagesService) GetMessage(id string) (*model.Message, error) {
	msg, err := s.cashRepo.Read(id)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
