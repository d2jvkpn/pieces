```bash
for i in {1..3}; do
    d="$PWD/data/mongodb-replSetI-${i}"
    mkdir -p $d/db
    mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI &
done
```

    [1]   Done                    mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI
    [2]   Done                    mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI
    [3]   Done                    mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI
    [4]-  Done                    mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI
    [5]+  Done                    mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI
    [1] 26496
    [2] 26498
    [3] 26500



```bash
netstat -nltp | grep mongod
```

    (Not all processes could be identified, non-owned process info
     will not be shown, you would have to be root to see it all.)
    tcp        0      0 127.0.0.1:2701          0.0.0.0:*               LISTEN      26496/[01;31m[Kmongod[m[K        
    tcp        0      0 127.0.0.1:2702          0.0.0.0:*               LISTEN      26498/[01;31m[Kmongod[m[K        
    tcp        0      0 127.0.0.1:2703          0.0.0.0:*               LISTEN      26500/[01;31m[Kmongod[m[K        



```bash
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
```

    MongoDB shell version v3.6.3
    connecting to: mongodb://localhost:2701/
    MongoDB server version: 3.6.3
    {
    	"ok" : 1,
    	"operationTime" : Timestamp(1603781378, 1),
    	"$clusterTime" : {
    		"clusterTime" : Timestamp(1603781378, 1),
    		"signature" : {
    			"hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
    			"keyId" : NumberLong(0)
    		}
    	}
    }
    bye



```bash
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
```

    >>> master node: localhost:2701
    MongoDB shell version v3.6.3
    connecting to: mongodb://localhost:2701/
    MongoDB server version: 3.6.3
    {
    	"set" : "replSetI",
    	"date" : ISODate("2020-10-27T06:49:54.279Z"),
    	"myState" : 1,
    	"term" : NumberLong(1),
    	"heartbeatIntervalMillis" : NumberLong(2000),
    	"optimes" : {
    		"lastCommittedOpTime" : {
    			"ts" : Timestamp(1603781390, 5),
    			"t" : NumberLong(1)
    		},
    		"readConcernMajorityOpTime" : {
    			"ts" : Timestamp(1603781390, 5),
    			"t" : NumberLong(1)
    		},
    		"appliedOpTime" : {
    			"ts" : Timestamp(1603781390, 5),
    			"t" : NumberLong(1)
    		},
    		"durableOpTime" : {
    			"ts" : Timestamp(1603781390, 5),
    			"t" : NumberLong(1)
    		}
    	},
    	"members" : [
    		{
    			"_id" : 1,
    			"name" : "localhost:2701",
    			"health" : 1,
    			"state" : 1,
    			"stateStr" : "PRIMARY",
    			"uptime" : 29,
    			"optime" : {
    				"ts" : Timestamp(1603781390, 5),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-27T06:49:50Z"),
    			"infoMessage" : "could not find member to sync from",
    			"electionTime" : Timestamp(1603781388, 1),
    			"electionDate" : ISODate("2020-10-27T06:49:48Z"),
    			"configVersion" : 1,
    			"self" : true
    		},
    		{
    			"_id" : 2,
    			"name" : "localhost:2702",
    			"health" : 1,
    			"state" : 2,
    			"stateStr" : "SECONDARY",
    			"uptime" : 15,
    			"optime" : {
    				"ts" : Timestamp(1603781390, 5),
    				"t" : NumberLong(1)
    			},
    			"optimeDurable" : {
    				"ts" : Timestamp(1603781390, 5),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-27T06:49:50Z"),
    			"optimeDurableDate" : ISODate("2020-10-27T06:49:50Z"),
    			"lastHeartbeat" : ISODate("2020-10-27T06:49:52.930Z"),
    			"lastHeartbeatRecv" : ISODate("2020-10-27T06:49:50.762Z"),
    			"pingMs" : NumberLong(0),
    			"syncingTo" : "localhost:2701",
    			"configVersion" : 1
    		},
    		{
    			"_id" : 3,
    			"name" : "localhost:2703",
    			"health" : 1,
    			"state" : 2,
    			"stateStr" : "SECONDARY",
    			"uptime" : 15,
    			"optime" : {
    				"ts" : Timestamp(1603781390, 5),
    				"t" : NumberLong(1)
    			},
    			"optimeDurable" : {
    				"ts" : Timestamp(1603781390, 5),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-27T06:49:50Z"),
    			"optimeDurableDate" : ISODate("2020-10-27T06:49:50Z"),
    			"lastHeartbeat" : ISODate("2020-10-27T06:49:52.930Z"),
    			"lastHeartbeatRecv" : ISODate("2020-10-27T06:49:50.756Z"),
    			"pingMs" : NumberLong(0),
    			"syncingTo" : "localhost:2701",
    			"configVersion" : 1
    		}
    	],
    	"ok" : 1,
    	"operationTime" : Timestamp(1603781390, 5),
    	"$clusterTime" : {
    		"clusterTime" : Timestamp(1603781390, 5),
    		"signature" : {
    			"hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
    			"keyId" : NumberLong(0)
    		}
    	}
    }
    bye



```bash
i=4
d="$PWD/data/mongodb-replSetI-${i}"
mkdir -p $d/db
mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI &
echo "rs.add(\"localhost:270${i}\")" | mongo --host $master
```

    [4] 26786
    MongoDB shell version v3.6.3
    connecting to: mongodb://localhost:2701/
    MongoDB server version: 3.6.3
    {
    	"ok" : 1,
    	"operationTime" : Timestamp(1603781399, 1),
    	"$clusterTime" : {
    		"clusterTime" : Timestamp(1603781399, 1),
    		"signature" : {
    			"hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
    			"keyId" : NumberLong(0)
    		}
    	}
    }
    bye



```bash
i=5
d="$PWD/data/mongodb-replSetI-arb"
mkdir -p $d/db
mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI &
echo "rs.addArb(\"localhost:270${i} \")" | mongo --host $master

echo "rs.status();" | mongo --host $master
```

    [5] 26894
    MongoDB shell version v3.6.3
    connecting to: mongodb://localhost:2701/
    MongoDB server version: 3.6.3
    {
    	"ok" : 1,
    	"operationTime" : Timestamp(1603781404, 1),
    	"$clusterTime" : {
    		"clusterTime" : Timestamp(1603781404, 1),
    		"signature" : {
    			"hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
    			"keyId" : NumberLong(0)
    		}
    	}
    }
    bye
    MongoDB shell version v3.6.3
    connecting to: mongodb://localhost:2701/
    MongoDB server version: 3.6.3
    {
    	"set" : "replSetI",
    	"date" : ISODate("2020-10-27T06:50:05.060Z"),
    	"myState" : 1,
    	"term" : NumberLong(1),
    	"heartbeatIntervalMillis" : NumberLong(2000),
    	"optimes" : {
    		"lastCommittedOpTime" : {
    			"ts" : Timestamp(1603781404, 1),
    			"t" : NumberLong(1)
    		},
    		"readConcernMajorityOpTime" : {
    			"ts" : Timestamp(1603781404, 1),
    			"t" : NumberLong(1)
    		},
    		"appliedOpTime" : {
    			"ts" : Timestamp(1603781404, 1),
    			"t" : NumberLong(1)
    		},
    		"durableOpTime" : {
    			"ts" : Timestamp(1603781404, 1),
    			"t" : NumberLong(1)
    		}
    	},
    	"members" : [
    		{
    			"_id" : 1,
    			"name" : "localhost:2701",
    			"health" : 1,
    			"state" : 1,
    			"stateStr" : "PRIMARY",
    			"uptime" : 40,
    			"optime" : {
    				"ts" : Timestamp(1603781404, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-27T06:50:04Z"),
    			"infoMessage" : "could not find member to sync from",
    			"electionTime" : Timestamp(1603781388, 1),
    			"electionDate" : ISODate("2020-10-27T06:49:48Z"),
    			"configVersion" : 3,
    			"self" : true
    		},
    		{
    			"_id" : 2,
    			"name" : "localhost:2702",
    			"health" : 1,
    			"state" : 2,
    			"stateStr" : "SECONDARY",
    			"uptime" : 26,
    			"optime" : {
    				"ts" : Timestamp(1603781399, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDurable" : {
    				"ts" : Timestamp(1603781399, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-27T06:49:59Z"),
    			"optimeDurableDate" : ISODate("2020-10-27T06:49:59Z"),
    			"lastHeartbeat" : ISODate("2020-10-27T06:50:04.918Z"),
    			"lastHeartbeatRecv" : ISODate("2020-10-27T06:50:04.920Z"),
    			"pingMs" : NumberLong(0),
    			"syncingTo" : "localhost:2701",
    			"configVersion" : 2
    		},
    		{
    			"_id" : 3,
    			"name" : "localhost:2703",
    			"health" : 1,
    			"state" : 2,
    			"stateStr" : "SECONDARY",
    			"uptime" : 26,
    			"optime" : {
    				"ts" : Timestamp(1603781399, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDurable" : {
    				"ts" : Timestamp(1603781399, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-27T06:49:59Z"),
    			"optimeDurableDate" : ISODate("2020-10-27T06:49:59Z"),
    			"lastHeartbeat" : ISODate("2020-10-27T06:50:04.918Z"),
    			"lastHeartbeatRecv" : ISODate("2020-10-27T06:50:04.920Z"),
    			"pingMs" : NumberLong(0),
    			"syncingTo" : "localhost:2701",
    			"configVersion" : 2
    		},
    		{
    			"_id" : 4,
    			"name" : "localhost:2704",
    			"health" : 1,
    			"state" : 2,
    			"stateStr" : "SECONDARY",
    			"uptime" : 3,
    			"optime" : {
    				"ts" : Timestamp(1603781399, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDurable" : {
    				"ts" : Timestamp(1603781399, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-27T06:49:59Z"),
    			"optimeDurableDate" : ISODate("2020-10-27T06:49:59Z"),
    			"lastHeartbeat" : ISODate("2020-10-27T06:50:04.918Z"),
    			"lastHeartbeatRecv" : ISODate("2020-10-27T06:50:04.920Z"),
    			"pingMs" : NumberLong(0),
    			"configVersion" : 2
    		},
    		{
    			"_id" : 5,
    			"name" : "localhost:2705",
    			"health" : 0,
    			"state" : 8,
    			"stateStr" : "(not reachable/healthy)",
    			"uptime" : 0,
    			"lastHeartbeat" : ISODate("2020-10-27T06:50:04.919Z"),
    			"lastHeartbeatRecv" : ISODate("1970-01-01T00:00:00Z"),
    			"pingMs" : NumberLong(0),
    			"lastHeartbeatMessage" : "Connection refused",
    			"configVersion" : -1
    		}
    	],
    	"ok" : 1,
    	"operationTime" : Timestamp(1603781404, 1),
    	"$clusterTime" : {
    		"clusterTime" : Timestamp(1603781404, 1),
    		"signature" : {
    			"hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
    			"keyId" : NumberLong(0)
    		}
    	}
    }
    bye



```bash
jobs
```

    [1]   Running                 mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI &
    [2]   Running                 mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI &
    [3]   Running                 mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI &
    [4]-  Running                 mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI &
    [5]+  Running                 mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI &



```bash
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
```

    MongoDB shell version v3.6.3
    connecting to: mongodb://localhost:2701/
    MongoDB server version: 3.6.3
    switched to db replSetI
    BulkWriteResult({
    	"writeErrors" : [ ],
    	"writeConcernErrors" : [ ],
    	"nInserted" : 2,
    	"nUpserted" : 0,
    	"nMatched" : 0,
    	"nModified" : 0,
    	"nRemoved" : 0,
    	"upserted" : [ ]
    })
    {
    	"_id" : ObjectId("5f97c32e20a3f32dd5f706a6"),
    	"title" : "MongoDB Overview",
    	"description" : "MongoDB is no SQL database",
    	"by" : "tutorials point",
    	"url" : "http://www.tutorialspoint.com",
    	"tags" : [
    		"mongodb",
    		"database",
    		"NoSQL"
    	],
    	"likes" : 100
    }
    {
    	"_id" : ObjectId("5f97c32e20a3f32dd5f706a7"),
    	"title" : "NoSQL Database",
    	"description" : "NoSQL database doesn't have tables",
    	"by" : "tutorials point",
    	"url" : "http://www.tutorialspoint.com",
    	"tags" : [
    		"mongodb",
    		"database",
    		"NoSQL"
    	],
    	"likes" : 20,
    	"comments" : [
    		{
    			"user" : "user1",
    			"message" : "My first comment",
    			"dateCreated" : ISODate("2013-12-09T18:35:00Z"),
    			"like" : 0
    		}
    	]
    }
    true
    bye



```bash
kill %1
```


```bash
jobs
```

    [1]   Done                    mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI
    [2]   Running                 mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI &
    [3]   Running                 mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI &
    [4]-  Running                 mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI &
    [5]+  Running                 mongod --port 270${i} --dbpath="$d/db" -logpath="$d/mongo.log" --replSet replSetI &



```bash

```


```bash
# kill %{1..5}
```
