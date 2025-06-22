package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("iniciando o cliente TCP")

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("erro ao conectar ao servidor:", err)
		return
	}
	defer conn.Close()

	fmt.Println("conectado ao servidor em localhost:8000")

	message := "oi servidor, sou um cliente"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("erro ao enviar dados para o servidor:", err)
		return
	}
	fmt.Println("mensagem enviada para o servidor")

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("erro ao ler dados do servidor:", err)
		return
	}

	response := string(buffer[:n])
	fmt.Printf("resposta recebida do servidor: %s\n", response)
}
