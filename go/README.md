# Go + Mongo

## Install Go

```console
export GO_VERSION=1.13.6
curl -sfLo go$GO_VERSION.linux-amd64.tar.gz https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz

nano ~/.profile
--add the following line to the bottom of the file
export PATH=$PATH:/usr/local/go/bin
-- save and exit
source $HOME/.profile
go version
```

## Connect

```console
cd src
go get go.mongodb.org/mongo-driver/mongo
go run connect.go
```