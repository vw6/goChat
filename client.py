#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import socket, curses
from threading import Thread

def chat():
    # myscreen.addstr(1, 1, "Ваше сообщение: ")
    # mess = bytearray(myscreen.getstr(1, 17))
    # sock.send(mess)
    myscreen.getch()

def respMess():
    while 1 :
        data = sock.recv(1024).decode('utf-8')
        myscreen.addstr(2, 1, data)


try: 
    myscreen = curses.initscr()
    myscreen.border(0)
    sock = socket.socket()
    sock.connect(('localhost', 8080))

    t1 = Thread(target=respMess)
    t1.start()
    t1.join()

    chat()
finally: 
    sock.close()
    curses.endwin()
