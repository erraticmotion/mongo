# Install Mongo Ruby 
 
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

## File Watcher
Visual Studio Code is unable to watch for file changes in this large workspace (error ENOSPC)

cat /proc/sys/fs/inotify/max_user_watches
8192
sudo nano /etc/sysctl.conf

-- add at the end of the file
fs.inotify.max_user_watches=524288
-- save and close

sudo sysctl -p
cat /proc/sys/fs/inotify/max_user_watches
524288

## VS Code Extensions
- Docker
- Docker Explorer
- Ruby langauge support