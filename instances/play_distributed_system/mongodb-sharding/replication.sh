#! /usr/bin/env bash
set -eu -o pipefail

# mkdir -p data/mongodb-replSetI-{1..3}/db

#### build cluster
for i in {1..3}; do
    d="$PWD/data/mongodb-replSetI-${i}"
    mkdir -p $d/db
    mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI &
done

netstat -nltp | grep mongod

cat > replSetI_config.js <<EOF
var cfg = {
  '_id':'replSetI',
  'members':[
    {'_id':1, 'host': 'localhost:2701'},
    {'_id':2, 'host': 'localhost:2702'},
    {'_id':3, 'host': 'localhost:2703'},
  ]
};

rs.initiate(cfg);
EOF

mongo --host localhost:2701 < replSetI_config.js
rm replSetI_config.js

#### check primary
master=""
for i in {1..3}; do
    host="localhost:270$i"
    ismaster=$(echo "rs.isMaster()['ismaster']" | mongo --quiet --host $host)

    if [[ "$ismaster" == "true" ]]; then
        master=$host
        echo ">>> master node: "$master
        break
    fi
done

echo "rs.status();" | mongo --host $master


#### add secondary
i=4
d="$PWD/data/mongodb-replSetI-${i}"
mkdir -p $d/db
mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI &
echo "rs.add(\"localhost:270${i} \")" | mongo --host $master


#### add arbiter
i=5
d="$PWD/data/mongodb-replSetI-arb"
mkdir -p $d/db
mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI &
echo "rs.addArb(\"localhost:270${i} \")" | mongo --host $master

echo "rs.status();" | mongo --host $master


#### test
cat > tmp.js << EOF
use replSetI;

db.mycol.insert([
  {
    title: "MongoDB Overview",
    description: "MongoDB is no SQL database",
    by: "tutorials point",
    url: "http://www.tutorialspoint.com",
    tags: ["mongodb", "database", "NoSQL"],
    likes: 100
  },
  {
    title: "NoSQL Database",
    description: "NoSQL database doesn't have tables",
    by: "tutorials point",
    url: "http://www.tutorialspoint.com",
    tags: ["mongodb", "database", "NoSQL"],
    likes: 20,
    comments: [
      {
        user:"user1",
        message: "My first comment",
        dateCreated: new Date(2013,11,10,2,35),
        like: 0
      }
    ]
  }
]);

db.mycol.find().pretty();

db.mycol.drop();
EOF

mongo --host $master < tmp.js
rm tmp.js

wait
