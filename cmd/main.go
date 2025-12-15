package main

import (
	"fmt"
	server "taskboard/internal"
)

func main() {
	serv, err := server.NewServer()
	if err != nil {
		panic(err)
	}
	urlApp := fmt.Sprintf("%s:%d", serv.Config().Server.Host, serv.Config().Server.Port)
	if err := serv.App().Listen(urlApp); err != nil {
		serv.Logger().Fatal("", err)
	}
	// TODO: запустить gRPC-сервер приложения
}
