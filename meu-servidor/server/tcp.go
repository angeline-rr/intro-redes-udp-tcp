package server

import (
	"fmt"
	"net"
)

func IniciarServidor() {
	ln, _ := net.Listen("tcp", ":8080")
	fmt.Println("Servidor ouvindo na porta 8080")

	for {
		conn, _ := ln.Accept()
		go lidarComConexao(conn)
	}
}

func lidarComConexao(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	fmt.Println("Recebido do cliente:", string(buffer[:n]))
	conn.Write([]byte("Olá do servidor TCP\n"))
}
