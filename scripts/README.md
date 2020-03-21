# Scripts

```console
sudo systemctl status mongodb

mongo --eval 'db.runCommand({ connectionStatus: 1 })'

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

Connect to the `tutorial` Mongo database

```console
irb -r ./connect.rb
exit
```
-or-

```console
ruby connect.rb
```