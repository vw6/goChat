#!/usr/bin/env python
# -*- coding: utf-8 -*-

import socket

sock = socket.socket()
sock.connect(('localhost', 8080))
mess = bytearray(input('Ваше сообщение серверу: ').encode('utf-8'))
sock.send(mess)

# data = sock.recv(1024).decode('utf-8')
sock.close()

# print(data)
