#!/usr/bin/env bash

# Ref: https://github.com/microservices-demo/user/blob/master/docker/user-db/scripts/mongo_create_insert.sh

SCRIPT_DIR=$(dirname "$0")

mongod --fork --logpath /data/db-users/mongodb.log --dbpath /data/db-users/
echo $?

# delay while starting up
echo "Delaying startup... 30s"
sleep 30s
/usr/bin/mongo --eval "db.version()"

echo "Preparing to run scripts"
# Create the Accounts
mongo localhost:27017/users accounts-create.js

# Insert useful data
mongo localhost:27017/users address-insert.js
mongo localhost:27017/users card-insert.js
mongo localhost:27017/users customer-insert.js

mongod --dbpath /data/db-users --shutdown