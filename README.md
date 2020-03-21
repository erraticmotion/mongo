# Mongo

Mongo, Ruby and Go development

## Install Mongo BD

```console
wget -qO - https://www.mongodb.org/static/pgp/server-4.2.asc | sudo apt-key add -
echo "deb [ arch=amd64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.2.list
sudo apt-get update
sudo apt install -y mongodb
```

## Install Ruby

```console
sudo apt-get install -y ruby
sudo apt-get install -y ruby-dev
sudo apt-get install -y build-essential
ruby --version
gem --version
sudo gem install mongo
```

Mongo DB needs a data directory. By default, the `mongod' daemon will store its data files in /data/db/. Create directory, 
and ensure it has the proper permissions:
```console
sudo mkdir -p /data/db/
sudo chown `id -u` /data/db
```

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

## Install direnv

```console
export DE_VERSION=2.20.0
curl -sfLo direnv https://github.com/direnv/direnv/releases/download/v$DE_VERSION/direnv.linux-amd64
chmod +x direnv
sudo mv direnv /usr/local/bin

nano ~/.bashrc
--add the following line to the bottom of the file
eval "$(direnv hook bash)"
-- save and exit
exec bash
```