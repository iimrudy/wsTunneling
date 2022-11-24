"# wsTunneling" 
this project is ment to bypass my schools internet limitations.

basically my school blocks all protocols except for http, https, SMTP and POP3. so i made this project to bypass that.

the app is ment to create a websocket server that will tunnel all the traffic generated from a specified tcp connection to a websocket client.


## Server mode


the sh ip the ssh server that you want to tunnel
the ip and port is the port of the websocket server (http servers)
```ssh
./sshbypass -mode server -ssh-ip localhost -ssh-port 22 -ip 0.0.0.0 -port 8080
```

## Client mode
the will be the new tcp server that will be created and where you will connect to
the ip and port is the port where the websocket server is running
```ssh
./sshbypass -mode client -ssh-ip localhost -ssh-port 22 -ip 1.1.1.1 -port 8080
```

