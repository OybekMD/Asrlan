# Install tools

Go programming language

```sh
$ wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz

$ tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz

$ export PATH=$PATH:/usr/local/go/bin
$ export GOPATH=$HOME/go
$ export PATH=$PATH:$GOPATH/bin

$ go version
```

Redis Database for cashe

```sh
$ sudo apt install lsb-release curl gpg

$ curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg

$ echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list

$ sudo apt-get update

$ sudo apt-get install redis

$ redis-cli
```

# Running app

```sh
$ go run cmd/main.go
```
or 
```sh
$ make run
```
# Asrlan
