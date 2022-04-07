package main

import (
	"fmt"
	"os"
	"os/signal"
	"wbl0/pkg/handler"
	"wbl0/pkg/model"
	"wbl0/pkg/nats"
	"wbl0/pkg/repository"
	"wbl0/pkg/service"
)

func main() {
	msgs := make(map[string]*model.Message)

	r, err := repository.NewRepository(msgs)
	if err != nil {
		panic(err)
	}

	s := service.NewService(r)

	err = s.RestoreCash()
	if err != nil {
		panic(err)
	}

	h := handler.NewHandler(s)

	go h.StartServer()

	n := nats.NewNats(s)

	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, os.Interrupt)

	fmt.Println("Nats is run")
	n.RunSubscribe(signalChan)
}
