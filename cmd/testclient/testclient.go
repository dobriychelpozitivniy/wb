package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"wbl0/pkg/model"

	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", "wwwwqqqwwwwwwww")
	if err != nil {
		panic(fmt.Sprintf("Error connect to nats: %s", err))
	}

	m := readModel()

	msg, err := json.Marshal(m)
	if err != nil {
		fmt.Errorf("Error marshal struct from file to json: %s", err)
	}

	err = sc.Publish("xxxxx", msg)
	if err != nil {
		panic(fmt.Sprintf("Error publish to nats: %s", err))
	}
}

func readModel() *model.Message {
	jsonFile, err := os.Open("model.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")

	defer jsonFile.Close()

	var message model.Message
	file, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal([]byte(file), &message)
	if err != nil {
		fmt.Errorf("Error unmarshal: %s", err)
	}

	id := uuid.New().String()
	message.OrderUID = id

	fmt.Println(message)

	return &message
}
