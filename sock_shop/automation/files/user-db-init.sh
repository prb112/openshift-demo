#!/usr/bin/env bash

# Ref: https://github.com/microservices-demo/user/blob/master/docker/user-db/scripts/mongo_create_insert.sh

SCRIPT_DIR=$(dirname "$0")

mongod --fork --logpath /var/log/mongodb.log --dbpath /data/db/

# Ref: https://stackoverflow.com/questions/26558932/what-is-the-correct-way-to-wait-until-mongodb-is-ready-after-restart
while ! /usr/bin/mongo --eval "db.version()" > /dev/null 2>&1
do
    sleep 0.1
done

FILES=$SCRIPT_DIR/*-create.js
for f in $FILES; do mongo localhost:27017/users $f; done

FILES=$SCRIPT_DIR/*-insert.js
for f in $FILES; do mongo localhost:27017/users $f; done