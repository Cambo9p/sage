import socket


def send_to_server(message):
    # TODO import params from yml file
    HOST = 'localhost'
    PORT = 5000

    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.connect((HOST, PORT))
        s.sendall(message.encode())

    return
