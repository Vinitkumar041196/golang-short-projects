# Welcome to README of pingcli

Pingcli is a demo project showcasing Golang's ability to create CLI utilities and execute os commands programatically.

I have created a simple CLI program that keeps on accepting user input using bufio Scanner from STDIN and execute the command.
It supports two commands: 
1. "exit" or "q" or "quit" : to exit program
2. "ping": takes two args
    - IP address or url to ping 
    - Count or number of packet to send for ping

## Sample execution:

Command:

```
go run cli.go
```

![console with success and error output](image.png)