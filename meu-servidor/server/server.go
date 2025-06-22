package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("iniciando o servidor TCP")

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("erro ao iniciar o servidor:", err)
		return
	}
	defer listener.Close()

	fmt.Println("servidor escutando em localhost:8000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("erro ao aceitar conexão:", err)
			continue
		}

		fmt.Println("cliente conectado de:", conn.RemoteAddr().String())

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("erro ao ler dados do cliente:", err)
		return
	}

	message := string(buffer[:n])
	fmt.Printf("mensagem recebida do cliente: %s\n", message)

	response := "oi cliente, mensagem recebida"
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("erro ao escrever para o cliente:", err)
		return
	}
	fmt.Println("resposta enviada")
}
