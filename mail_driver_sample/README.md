## installation
clone project
```
$ cd &GOPATH/src/github.com
$ mkdir go-examples
$ cd go-examples
$ git clone github.com/y-ogura/go-examples
$ dep ensure
$ cp .envrc.example .envrc
```

edit .envrc
if use sparkpost
```
export MAIL_DRIVER="sparkpost"
export SPARKPOST_API_KEY="your api key"
```

run
```
$ go run main.go
```
