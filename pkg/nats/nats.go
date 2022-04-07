package nats

import (
	"encoding/json"
	"fmt"
	"os"
	"wbl0/pkg/model"
	"wbl0/pkg/service"

	"github.com/jackc/pgx/v4"
	stan "github.com/nats-io/stan.go"
	"gopkg.in/validator.v2"
)

type Nats struct {
	service *service.Service
}

func NewNats(s *service.Service) *Nats {
	return &Nats{service: s}
}

func (n *Nats) RunSubscribe(exit <-chan os.Signal) {
	sc, err := stan.Connect("test-cluster", "clientid")
	if err != nil {
		panic(fmt.Sprintf("Error connect to nats: %s", err))
	}

	startOpt := stan.DeliverAllAvailable()

	t, err := n.service.GetLastTime()
	if err != nil && err != pgx.ErrNoRows {
		panic(fmt.Sprintf("Error get last time: %s", err))
	}

	if err == nil && err != pgx.ErrNoRows {
		fmt.Println("start at time")
		startOpt = stan.StartAtTime(*t)
	}

	sub, err := sc.Subscribe("xxxxx", func(m *stan.Msg) {
		msg, err := bytesToStruct(m.Data)
		if err != nil {
			fmt.Printf("Error unmarshal message from nats: %s", err)
			m.Ack()

			return
		}

		err = n.service.AddMessage(msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(msg, "is added")

		m.Ack()
	}, stan.DurableName("durable"), stan.SetManualAckMode(), startOpt)

	if err != nil {
		panic(fmt.Sprintf("Error subscribe to nats: %s", err))
	}

	cleanupDone := make(chan bool)
	go func() {
		for range exit {
			fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			sub.Unsubscribe()
			sc.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}

func bytesToStruct(b []byte) (*model.Message, error) {
	var msg model.Message

	err := json.Unmarshal(b, &msg)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshal message to struct: %s", err)
	}

	if errs := validator.Validate(msg); errs != nil {
		return nil, fmt.Errorf("Error validate message: %s", errs)
	}

	return &msg, nil
}
