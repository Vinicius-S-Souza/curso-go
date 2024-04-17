package main

import (
	"fmt"

	"github.com/devfullcycle/fcutils/pkg/rabbitmq"
)

func main() {

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	var msg string

	for i := 0; i < 100000; i++ {

		msg = fmt.Sprintf("Mensagem %d", i)

		rabbitmq.Publish(ch, msg, "amq.direct")
	}
	

}
