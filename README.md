# api-logs
Save logs from angular app


### Build
```
$ env GOOS=linux GOARCH=amd64 go build -o logs-linux-amd64 cmd/api/main.go
```

### Using Shell Script
```shell
$ ./go-executable-build.bash -o api_logs -p cmd/api/main.go
# select one architecture
$ ./go-executable-build.bash -o api_logs -p cmd/api/main.go -a linux/amd64
```


####
|GOOS - Target Operating System	|GOARCH - Target Platform|
| ------ | ------ |
|android|	arm|
|darwin|	386|
|darwin|	amd64|
|darwin|	arm|
|darwin|	arm64|
|dragonfly|	amd64|
|freebsd|	386|
|freebsd|	amd64|
|freebsd|	arm|
|linux|	386|
|linux|	amd64|
|linux|	arm|
|linux|	arm64|
|linux|	ppc64|
|linux|	ppc64le|
|linux|	mips|
|linux|	mipsle|
|linux|	mips64|
|linux|	mips64le|
|netbsd|	386|
|netbsd|	amd64|
|netbsd|	arm|
|openbsd|	386|
|openbsd|	amd64|
|openbsd|	arm|
|plan9|	386|
|plan9|	amd64|
|solaris|	amd64|
|windows|	386|
|windows|	amd64|

### Create a Daemon

```shell
$ sudo touch /lib/systemd/system/apilog.service
```
Edit this file
```shell
$ sudo nano /lib/systemd/system/apilog.service
```

put this into a file

```shell
[Unit]
Description=apilog

[Service]
Type=simple
Restart=always
RestartSec=30s
Environment=PORT=8080
ExecStart=/path/to/file/apilog-linux-amd64

[Install]
WantedBy=multi-user.target
```

Now you can use this commands
```shell
$ systemctl list-units
$ systemctl is-active apilog.service
$ systemctl is-enabled apilog.service
$ systemctl is-failed apilog.service
$ sudo systemctl start apilog.service
$ sudo systemctl stop apilog.service
$ sudo systemctl enable apilog.service
$ sudo systemctl disable apilog.service
$ sudo systemctl status apilog.service
```