import socket


def send_message(host: str, port: int, message: str):
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

    server_socket.sendto(message, (host, port))


if __name__ == '__main__':
    HOST = 'localhost'
    PORT = 9000

    while True:
        message = input('Mensagem: ').encode('utf-8')
        send_message(HOST, PORT, message)
