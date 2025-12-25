package main

import (
	"fmt"
	"taskboard/internal/infrastructure/gateways/http"
)

func main() {
	serv, err := http.NewServer()
	if err != nil {
		panic(err)
	}
	log := serv.Logger()
	urlApp := fmt.Sprintf("%s:%d", serv.Config().Server.Host, serv.Config().Server.Port)
	if err := serv.App().Listen(urlApp); err != nil {
		log.Fatal("Error with init server", err)
	}
	// TODO: запустить gRPC-сервер приложения
}
