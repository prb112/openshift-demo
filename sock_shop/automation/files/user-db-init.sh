#!/usr/bin/env bash

# Ref: https://github.com/microservices-demo/user/blob/master/docker/user-db/scripts/mongo_create_insert.sh

SCRIPT_DIR=$(dirname "$0")

mongod --fork --logpath /data/db-users/mongodb.log --dbpath /data/db-users/

# Ref: https://stackoverflow.com/questions/26558932/what-is-the-correct-way-to-wait-until-mongodb-is-ready-after-restart
# delay while starting up
while ! /usr/bin/mongo --eval "db.version()" > /dev/null 2>&1; do sleep 0.1; done

# Create the Accounts
mongo localhost:27017/users accounts-create.js

# Insert useful data
mongo localhost:27017/users address-insert.js
mongo localhost:27017/users card-insert.js
mongo localhost:27017/users customer-insert.js
