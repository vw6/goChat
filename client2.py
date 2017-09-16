#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import socket, curses

sock = socket.socket()
sock.connect(('localhost', 8080))

# myscreen = curses.initscr() 

data = sock.recv(1024).decode('utf-8')
print(data)

sock.close()