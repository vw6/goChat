import socket

sock = socket.socket()
sock.connect(('localhost', 80))
sock.setblocking(0)
mess = bytearray(input('Ваше сообщение серверу: ').encode('utf-8'))
sock.send(mess)

try: data = sock.recv(1024)
except socket.error:
    pass 
else: 
    print(data)




