# gologger

Sends messages to the local syslog daemon

Meant as a replacement for the default `logger` found on most linux systems ([docs](http://man7.org/linux/man-pages/man1/logger.1.html)). I built this to run on older installations, where the default `logger` has a hard limit for the size of each message (1024 bytes).

## Usage

As a pipe:

```bash
cat /some/file | gologger -t myfile
```

As arguments:

```bash
gologger -t myprog my message
gologger -t myprog "my   message    with weird          spaces"
```

In Apache (my use case):

```apache
ErrorLog "|/usr/bin/gologger -p local6.debug"
CustomLog "|/usr/bin/gologger -p local0.debug" combinedio env=!dontlog
```

## Compatibility

For now, only supports these flags:

* `-p`, `--priority`
* `-t`, `--tag`
