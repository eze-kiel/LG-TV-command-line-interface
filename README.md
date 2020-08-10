# CLI for LG television
Tool designed to send commands to a LG television via TCP/IP. Currently working with LG 55SL5B model.

### Usage
```
USAGE:
   LG-TV-command-line-interface [global options] command [command options] [arguments...]

COMMANDS:
   brightness  Adjusts the brightness of the screen. The value must be in 0-100
   mute        Turns on/off the TV sound. The argument must be true to mute, anything else to unmute
   poweroff    Turns off the TV. Needs no arguments
   volume      Adjusts the TV volume. The value must be in 0-100
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

### Example
```
$ lg-cli volume 50      //set the volume to 50
```

### Configuration
You have to specify the IP of your TV and the port used in the config.yaml file. By default, the port used should be 9761 or 9760.