package main

import "github.com/wandermaia/fc-utils/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	//amd.direct é uma fila default. É necessário fazer um bind da fila, no caso, a fila orders.
	rabbitmq.Publish(ch, "Hello World!", "amq.direct")
}
