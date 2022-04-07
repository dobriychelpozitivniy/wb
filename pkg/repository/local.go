package repository

import (
	"fmt"
	"sync"
	"wbl0/pkg/model"
)

type LocalRepository struct {
	Msgs  map[string]*model.Message
	mutex *sync.Mutex
}

func NewLocalRepository(msgs map[string]*model.Message) *LocalRepository {
	return &LocalRepository{Msgs: msgs, mutex: &sync.Mutex{}}
}

func (r *LocalRepository) Add(msg *model.Message) {
	r.mutex.Lock()
	r.Msgs[msg.OrderUID] = msg
	r.mutex.Unlock()
}

func (r *LocalRepository) Read(id string) (*model.Message, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	msg, ok := r.Msgs[id]
	if !ok {
		return nil, fmt.Errorf("Msg with id: %s not exist in local memory.", id)
	}

	return msg, nil
}

func (r *LocalRepository) AddMany(msgs []*model.Message) {
	r.mutex.Lock()
	for _, v := range msgs {
		r.Msgs[v.OrderUID] = v
	}
	r.mutex.Unlock()
}
