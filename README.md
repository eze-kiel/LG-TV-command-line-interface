# CLI for LG television
Tool designed to send commands to an LG television via TCP/IP. Currently working with LG 55SL5B model.
### Usage
For Linux :
```bash
GOOS=linux go build main.go -o cli-lg 
./cli-lg <param> <value>
```

For Windows :
```bash
GOOS=windows go build main.go -o cli-lg 
./cli-lg <param> <value>
```