# CLI for LG television
Tool designed to send commands to a LG television via TCP/IP. Currently working with LG 55SL5B model.
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

Example of use :
```bash
./cli-lg volume 50 #this will set the volume at 50
```

### Note
For now, you must have an environement variable 'LGIP' which contain the IP of the LG television you want to control.