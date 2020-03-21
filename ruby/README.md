# Ruby + Mongo

## Install Ruby

```console
sudo apt-get install -y ruby
sudo apt-get install -y ruby-dev
sudo apt-get install -y build-essential
ruby --version
gem --version
sudo gem install mongo
```

```console
mongo --eval 'db.runCommand({ connectionStatus: 1 })'
```

Connect to the `tutorial` Mongo database

```console
irb -r ./connect.rb
exit
```
-or-

```console
ruby connect.rb
```