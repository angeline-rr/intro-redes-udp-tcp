package client

import (
	"fmt"
	"net"
)

func IniciarCliente() {
	conn, _ := net.Dial("tcp", "localhost:8080")
	defer conn.Close()
	conn.Write([]byte("Mensagem do cliente!\n"))
	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	fmt.Println("Resposta do servidor:", string(buffer[:n]))
}
