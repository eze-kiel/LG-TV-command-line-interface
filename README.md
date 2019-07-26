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

### Configuration
You have to specify the IP of your TV in the config.yaml file.

### Functions
```bash
volume <value> #value between 0 and 100
mute <state> #boolean state
input <value> #HDMI1, RGB... Note that this field is case insensitive
poweroff #turn off the screen
```