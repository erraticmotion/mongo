# Mongo

Mongo development using Ruby and Go

## Install Mongo BD

```console
wget -qO - https://www.mongodb.org/static/pgp/server-4.2.asc | sudo apt-key add -
echo "deb [ arch=amd64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.2.list
sudo apt-get update
sudo apt install -y mongodb
```

Mongo DB needs a data directory. By default, the `mongod' daemon will store its data files in /data/db/. Create directory, 
and ensure it has the proper permissions:

```console
sudo mkdir -p /data/db/
sudo chown `id -u` /data/db
```

```console
sudo systemctl status mongodb
```

## Mongo DB shell

```console
mongo
> use tutorial

> db.users.insert({username: "smith"})

> db.users.find()

> db.users.insert({username: "jones"})

> db.users.count()

> db.users.find({username: "jones"})

> db.users.drop()

> exit
```