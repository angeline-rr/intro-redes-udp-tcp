package main

import (
	"time"

	"github.com/angeline-rr/meu-servidor/client"
	"github.com/angeline-rr/meu-servidor/server"
)

func main() {
	go server.IniciarServidor() // inicia o servidor em uma goroutine
	time.Sleep(1 * time.Second) // espera 1 segundo para o servidor iniciar
	client.IniciarCliente()     // conecta o cliente
}
