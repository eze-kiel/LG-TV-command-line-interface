# CLI for LG television
Tool designed to send commands to a LG television via TCP/IP. Currently working with LG 55SL5B model.

### Usage
For Linux :
```bash
GOOS=linux go build -o cli-lg 
./cli-lg <func> <value>
```

For Windows :
```bash
GOOS=windows go build -o cli-lg 
./cli-lg <func> <value>
```

Example of use :
```bash
./cli-lg volume 50 #this will set the volume at 50
```

### Configuration
You have to specify the IP of your TV and the port used in the config.yaml file. By default, the port used should be 9761 or 9760.

### Functions
```bash
volume <value> #value between 0 and 100
brightness <value> #value between 0 and 100
contrast <value> #value between 0 and 100
mute <state> #boolean state
input <value> #HDMI1, RGB... Note that this field is case insensitive
poweroff #turn off the screen
```