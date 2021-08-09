```bash
docker network create mongo-net
```

    5ef25379eaf91c443c378a0be652533bb823723c302b6f4df88b15c7d5c73359



```bash
mkdir -p Workpath

for i in {1..3}; do
    name="mongo-wk${i}"
    echo ">>> container "$name

    docker run --detach --name=$name --publish=270${i}:27017 \
    --volume=$PWD/Workpath:/mnt/Workpath --network=mongo-net\
    mongo mongod --replSet replSetI
done
```

    >>> container mongo-wk1
    dd383ac1459ddb4a99609ea9cdc6a4b6ebc0fc5085d8c5b8c971fae737950112
    >>> container mongo-wk2
    5ae93bde8e0d4bee67c4361dc84313b328a348b095dc750afbb6910d55599810
    >>> container mongo-wk3
    0dcd062a6da0f9aec5973adf9bd4cb35d8b8f3155e1b0bb5b40ab052cf0d20ed



```bash
docker ps
```

    CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                     NAMES
    0dcd062a6da0        mongo               "docker-entrypoint.s…"   2 seconds ago       Up 1 second         0.0.0.0:2703->27017/tcp   mongo-wk3
    5ae93bde8e0d        mongo               "docker-entrypoint.s…"   2 seconds ago       Up 2 seconds        0.0.0.0:2702->27017/tcp   mongo-wk2
    dd383ac1459d        mongo               "docker-entrypoint.s…"   3 seconds ago       Up 2 seconds        0.0.0.0:2701->27017/tcp   mongo-wk1



```bash
cat > Workpath/replSetI_config.js  <<EOF
var cfg = {
  '_id':'replSetI',
  'members':[
    {'_id':1, 'host': 'mongo-wk1:27017'},
    {'_id':2, 'host': 'mongo-wk2:27017'},
    {'_id':3, 'host': 'mongo-wk3:27017'},
  ]
};

rs.initiate(cfg);
EOF

 mongo --quiet --host=localhost:2701 < Workpath/replSetI_config.js
```

    {
    	"ok" : 1,
    	"$clusterTime" : {
    		"clusterTime" : Timestamp(1603852665, 1),
    		"signature" : {
    			"hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
    			"keyId" : NumberLong(0)
    		}
    	},
    	"operationTime" : Timestamp(1603852665, 1)
    }



```bash
echo "rs.status()" | mongo --quiet --host=localhost:2701
```

    {
    	"set" : "replSetI",
    	"date" : ISODate("2020-10-28T02:38:10.404Z"),
    	"myState" : 1,
    	"term" : NumberLong(1),
    	"syncSourceHost" : "",
    	"syncSourceId" : -1,
    	"heartbeatIntervalMillis" : NumberLong(2000),
    	"majorityVoteCount" : 2,
    	"writeMajorityCount" : 2,
    	"votingMembersCount" : 3,
    	"writableVotingMembersCount" : 3,
    	"optimes" : {
    		"lastCommittedOpTime" : {
    			"ts" : Timestamp(1603852677, 1),
    			"t" : NumberLong(1)
    		},
    		"lastCommittedWallTime" : ISODate("2020-10-28T02:37:57.079Z"),
    		"readConcernMajorityOpTime" : {
    			"ts" : Timestamp(1603852677, 1),
    			"t" : NumberLong(1)
    		},
    		"readConcernMajorityWallTime" : ISODate("2020-10-28T02:37:57.079Z"),
    		"appliedOpTime" : {
    			"ts" : Timestamp(1603852677, 1),
    			"t" : NumberLong(1)
    		},
    		"durableOpTime" : {
    			"ts" : Timestamp(1603852677, 1),
    			"t" : NumberLong(1)
    		},
    		"lastAppliedWallTime" : ISODate("2020-10-28T02:37:57.079Z"),
    		"lastDurableWallTime" : ISODate("2020-10-28T02:37:57.079Z")
    	},
    	"lastStableRecoveryTimestamp" : Timestamp(1603852675, 3),
    	"electionCandidateMetrics" : {
    		"lastElectionReason" : "electionTimeout",
    		"lastElectionDate" : ISODate("2020-10-28T02:37:55.505Z"),
    		"electionTerm" : NumberLong(1),
    		"lastCommittedOpTimeAtElection" : {
    			"ts" : Timestamp(0, 0),
    			"t" : NumberLong(-1)
    		},
    		"lastSeenOpTimeAtElection" : {
    			"ts" : Timestamp(1603852665, 1),
    			"t" : NumberLong(-1)
    		},
    		"numVotesNeeded" : 2,
    		"priorityAtElection" : 1,
    		"electionTimeoutMillis" : NumberLong(10000),
    		"numCatchUpOps" : NumberLong(0),
    		"newTermStartDate" : ISODate("2020-10-28T02:37:55.542Z"),
    		"wMajorityWriteAvailabilityDate" : ISODate("2020-10-28T02:37:57.013Z")
    	},
    	"members" : [
    		{
    			"_id" : 1,
    			"name" : "mongo-wk1:27017",
    			"health" : 1,
    			"state" : 1,
    			"stateStr" : "PRIMARY",
    			"uptime" : 62,
    			"optime" : {
    				"ts" : Timestamp(1603852677, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-28T02:37:57Z"),
    			"syncSourceHost" : "",
    			"syncSourceId" : -1,
    			"infoMessage" : "Could not find member to sync from",
    			"electionTime" : Timestamp(1603852675, 1),
    			"electionDate" : ISODate("2020-10-28T02:37:55Z"),
    			"configVersion" : 1,
    			"configTerm" : 1,
    			"self" : true,
    			"lastHeartbeatMessage" : ""
    		},
    		{
    			"_id" : 2,
    			"name" : "mongo-wk2:27017",
    			"health" : 1,
    			"state" : 2,
    			"stateStr" : "SECONDARY",
    			"uptime" : 25,
    			"optime" : {
    				"ts" : Timestamp(1603852677, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDurable" : {
    				"ts" : Timestamp(1603852677, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-28T02:37:57Z"),
    			"optimeDurableDate" : ISODate("2020-10-28T02:37:57Z"),
    			"lastHeartbeat" : ISODate("2020-10-28T02:38:09.522Z"),
    			"lastHeartbeatRecv" : ISODate("2020-10-28T02:38:09.033Z"),
    			"pingMs" : NumberLong(0),
    			"lastHeartbeatMessage" : "",
    			"syncSourceHost" : "mongo-wk1:27017",
    			"syncSourceId" : 1,
    			"infoMessage" : "",
    			"configVersion" : 1,
    			"configTerm" : 1
    		},
    		{
    			"_id" : 3,
    			"name" : "mongo-wk3:27017",
    			"health" : 1,
    			"state" : 2,
    			"stateStr" : "SECONDARY",
    			"uptime" : 25,
    			"optime" : {
    				"ts" : Timestamp(1603852677, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDurable" : {
    				"ts" : Timestamp(1603852677, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-28T02:37:57Z"),
    			"optimeDurableDate" : ISODate("2020-10-28T02:37:57Z"),
    			"lastHeartbeat" : ISODate("2020-10-28T02:38:09.521Z"),
    			"lastHeartbeatRecv" : ISODate("2020-10-28T02:38:09.032Z"),
    			"pingMs" : NumberLong(0),
    			"lastHeartbeatMessage" : "",
    			"syncSourceHost" : "mongo-wk1:27017",
    			"syncSourceId" : 1,
    			"infoMessage" : "",
    			"configVersion" : 1,
    			"configTerm" : 1
    		}
    	],
    	"ok" : 1,
    	"$clusterTime" : {
    		"clusterTime" : Timestamp(1603852677, 1),
    		"signature" : {
    			"hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
    			"keyId" : NumberLong(0)
    		}
    	},
    	"operationTime" : Timestamp(1603852677, 1)
    }



```bash

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
    MongoDB server version: 4.4.1
    WARNING: shell and server versions do not match
    {
    	"set" : "replSetI",
    	"date" : ISODate("2020-10-28T02:38:42.944Z"),
    	"myState" : 1,
    	"term" : NumberLong(1),
    	"syncSourceHost" : "",
    	"syncSourceId" : -1,
    	"heartbeatIntervalMillis" : NumberLong(2000),
    	"majorityVoteCount" : 2,
    	"writeMajorityCount" : 2,
    	"votingMembersCount" : 3,
    	"writableVotingMembersCount" : 3,
    	"optimes" : {
    		"lastCommittedOpTime" : {
    			"ts" : Timestamp(1603852715, 1),
    			"t" : NumberLong(1)
    		},
    		"lastCommittedWallTime" : ISODate("2020-10-28T02:38:35.567Z"),
    		"readConcernMajorityOpTime" : {
    			"ts" : Timestamp(1603852715, 1),
    			"t" : NumberLong(1)
    		},
    		"readConcernMajorityWallTime" : ISODate("2020-10-28T02:38:35.567Z"),
    		"appliedOpTime" : {
    			"ts" : Timestamp(1603852715, 1),
    			"t" : NumberLong(1)
    		},
    		"durableOpTime" : {
    			"ts" : Timestamp(1603852715, 1),
    			"t" : NumberLong(1)
    		},
    		"lastAppliedWallTime" : ISODate("2020-10-28T02:38:35.567Z"),
    		"lastDurableWallTime" : ISODate("2020-10-28T02:38:35.567Z")
    	},
    	"lastStableRecoveryTimestamp" : Timestamp(1603852675, 3),
    	"electionCandidateMetrics" : {
    		"lastElectionReason" : "electionTimeout",
    		"lastElectionDate" : ISODate("2020-10-28T02:37:55.505Z"),
    		"electionTerm" : NumberLong(1),
    		"lastCommittedOpTimeAtElection" : {
    			"ts" : Timestamp(0, 0),
    			"t" : NumberLong(-1)
    		},
    		"lastSeenOpTimeAtElection" : {
    			"ts" : Timestamp(1603852665, 1),
    			"t" : NumberLong(-1)
    		},
    		"numVotesNeeded" : 2,
    		"priorityAtElection" : 1,
    		"electionTimeoutMillis" : NumberLong(10000),
    		"numCatchUpOps" : NumberLong(0),
    		"newTermStartDate" : ISODate("2020-10-28T02:37:55.542Z"),
    		"wMajorityWriteAvailabilityDate" : ISODate("2020-10-28T02:37:57.013Z")
    	},
    	"members" : [
    		{
    			"_id" : 1,
    			"name" : "mongo-wk1:27017",
    			"health" : 1,
    			"state" : 1,
    			"stateStr" : "PRIMARY",
    			"uptime" : 94,
    			"optime" : {
    				"ts" : Timestamp(1603852715, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-28T02:38:35Z"),
    			"syncSourceHost" : "",
    			"syncSourceId" : -1,
    			"infoMessage" : "Could not find member to sync from",
    			"electionTime" : Timestamp(1603852675, 1),
    			"electionDate" : ISODate("2020-10-28T02:37:55Z"),
    			"configVersion" : 1,
    			"configTerm" : 1,
    			"self" : true,
    			"lastHeartbeatMessage" : ""
    		},
    		{
    			"_id" : 2,
    			"name" : "mongo-wk2:27017",
    			"health" : 1,
    			"state" : 2,
    			"stateStr" : "SECONDARY",
    			"uptime" : 57,
    			"optime" : {
    				"ts" : Timestamp(1603852715, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDurable" : {
    				"ts" : Timestamp(1603852715, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-28T02:38:35Z"),
    			"optimeDurableDate" : ISODate("2020-10-28T02:38:35Z"),
    			"lastHeartbeat" : ISODate("2020-10-28T02:38:41.531Z"),
    			"lastHeartbeatRecv" : ISODate("2020-10-28T02:38:41.038Z"),
    			"pingMs" : NumberLong(0),
    			"lastHeartbeatMessage" : "",
    			"syncSourceHost" : "mongo-wk1:27017",
    			"syncSourceId" : 1,
    			"infoMessage" : "",
    			"configVersion" : 1,
    			"configTerm" : 1
    		},
    		{
    			"_id" : 3,
    			"name" : "mongo-wk3:27017",
    			"health" : 1,
    			"state" : 2,
    			"stateStr" : "SECONDARY",
    			"uptime" : 57,
    			"optime" : {
    				"ts" : Timestamp(1603852715, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDurable" : {
    				"ts" : Timestamp(1603852715, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-28T02:38:35Z"),
    			"optimeDurableDate" : ISODate("2020-10-28T02:38:35Z"),
    			"lastHeartbeat" : ISODate("2020-10-28T02:38:41.530Z"),
    			"lastHeartbeatRecv" : ISODate("2020-10-28T02:38:41.037Z"),
    			"pingMs" : NumberLong(0),
    			"lastHeartbeatMessage" : "",
    			"syncSourceHost" : "mongo-wk1:27017",
    			"syncSourceId" : 1,
    			"infoMessage" : "",
    			"configVersion" : 1,
    			"configTerm" : 1
    		}
    	],
    	"ok" : 1,
    	"$clusterTime" : {
    		"clusterTime" : Timestamp(1603852715, 1),
    		"signature" : {
    			"hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
    			"keyId" : NumberLong(0)
    		}
    	},
    	"operationTime" : Timestamp(1603852715, 1)
    }
    bye



```bash
echo $master
```

    localhost:2701



```bash
i=4
name="mongo-wk${i}"
echo ">>> container "$name

docker run --detach --name=$name --publish=270${i}:27017 \
--volume=$PWD/Workpath:/mnt/Workpath --network=mongo-net \
mongo mongod --replSet replSetI

echo "rs.add(\"mongo-wk${i}:27017\")" | mongo --quiet --host $master
```

    >>> container mongo-wk4
    9270f1398a0ea406c784a902525651a03c011964f9e8f5775003dc31417217ee
    {
    	"operationTime" : Timestamp(1603852955, 1),
    	"ok" : 0,
    	"errmsg" : "Found two member configurations with same host field, members.3.host == members.4.host == mongo-wk4:27017",
    	"code" : 103,
    	"codeName" : "NewReplicaSetConfigurationIncompatible",
    	"$clusterTime" : {
    		"clusterTime" : Timestamp(1603852955, 1),
    		"signature" : {
    			"hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
    			"keyId" : NumberLong(0)
    		}
    	}
    }



```bash
echo "rs.status();" | mongo --quiet --host $master
```

    {
    	"set" : "replSetI",
    	"date" : ISODate("2020-10-28T02:43:07.260Z"),
    	"myState" : 1,
    	"term" : NumberLong(1),
    	"syncSourceHost" : "",
    	"syncSourceId" : -1,
    	"heartbeatIntervalMillis" : NumberLong(2000),
    	"majorityVoteCount" : 3,
    	"writeMajorityCount" : 3,
    	"votingMembersCount" : 4,
    	"writableVotingMembersCount" : 4,
    	"optimes" : {
    		"lastCommittedOpTime" : {
    			"ts" : Timestamp(1603852985, 1),
    			"t" : NumberLong(1)
    		},
    		"lastCommittedWallTime" : ISODate("2020-10-28T02:43:05.580Z"),
    		"readConcernMajorityOpTime" : {
    			"ts" : Timestamp(1603852985, 1),
    			"t" : NumberLong(1)
    		},
    		"readConcernMajorityWallTime" : ISODate("2020-10-28T02:43:05.580Z"),
    		"appliedOpTime" : {
    			"ts" : Timestamp(1603852985, 1),
    			"t" : NumberLong(1)
    		},
    		"durableOpTime" : {
    			"ts" : Timestamp(1603852985, 1),
    			"t" : NumberLong(1)
    		},
    		"lastAppliedWallTime" : ISODate("2020-10-28T02:43:05.580Z"),
    		"lastDurableWallTime" : ISODate("2020-10-28T02:43:05.580Z")
    	},
    	"lastStableRecoveryTimestamp" : Timestamp(1603852975, 1),
    	"electionCandidateMetrics" : {
    		"lastElectionReason" : "electionTimeout",
    		"lastElectionDate" : ISODate("2020-10-28T02:37:55.505Z"),
    		"electionTerm" : NumberLong(1),
    		"lastCommittedOpTimeAtElection" : {
    			"ts" : Timestamp(0, 0),
    			"t" : NumberLong(-1)
    		},
    		"lastSeenOpTimeAtElection" : {
    			"ts" : Timestamp(1603852665, 1),
    			"t" : NumberLong(-1)
    		},
    		"numVotesNeeded" : 2,
    		"priorityAtElection" : 1,
    		"electionTimeoutMillis" : NumberLong(10000),
    		"numCatchUpOps" : NumberLong(0),
    		"newTermStartDate" : ISODate("2020-10-28T02:37:55.542Z"),
    		"wMajorityWriteAvailabilityDate" : ISODate("2020-10-28T02:37:57.013Z")
    	},
    	"members" : [
    		{
    			"_id" : 1,
    			"name" : "mongo-wk1:27017",
    			"health" : 1,
    			"state" : 1,
    			"stateStr" : "PRIMARY",
    			"uptime" : 359,
    			"optime" : {
    				"ts" : Timestamp(1603852985, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-28T02:43:05Z"),
    			"syncSourceHost" : "",
    			"syncSourceId" : -1,
    			"infoMessage" : "",
    			"electionTime" : Timestamp(1603852675, 1),
    			"electionDate" : ISODate("2020-10-28T02:37:55Z"),
    			"configVersion" : 2,
    			"configTerm" : 1,
    			"self" : true,
    			"lastHeartbeatMessage" : ""
    		},
    		{
    			"_id" : 2,
    			"name" : "mongo-wk2:27017",
    			"health" : 1,
    			"state" : 2,
    			"stateStr" : "SECONDARY",
    			"uptime" : 322,
    			"optime" : {
    				"ts" : Timestamp(1603852985, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDurable" : {
    				"ts" : Timestamp(1603852985, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-28T02:43:05Z"),
    			"optimeDurableDate" : ISODate("2020-10-28T02:43:05Z"),
    			"lastHeartbeat" : ISODate("2020-10-28T02:43:05.805Z"),
    			"lastHeartbeatRecv" : ISODate("2020-10-28T02:43:05.812Z"),
    			"pingMs" : NumberLong(0),
    			"lastHeartbeatMessage" : "",
    			"syncSourceHost" : "mongo-wk1:27017",
    			"syncSourceId" : 1,
    			"infoMessage" : "",
    			"configVersion" : 2,
    			"configTerm" : 1
    		},
    		{
    			"_id" : 3,
    			"name" : "mongo-wk3:27017",
    			"health" : 1,
    			"state" : 2,
    			"stateStr" : "SECONDARY",
    			"uptime" : 322,
    			"optime" : {
    				"ts" : Timestamp(1603852985, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDurable" : {
    				"ts" : Timestamp(1603852985, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-28T02:43:05Z"),
    			"optimeDurableDate" : ISODate("2020-10-28T02:43:05Z"),
    			"lastHeartbeat" : ISODate("2020-10-28T02:43:05.805Z"),
    			"lastHeartbeatRecv" : ISODate("2020-10-28T02:43:05.811Z"),
    			"pingMs" : NumberLong(0),
    			"lastHeartbeatMessage" : "",
    			"syncSourceHost" : "mongo-wk1:27017",
    			"syncSourceId" : 1,
    			"infoMessage" : "",
    			"configVersion" : 2,
    			"configTerm" : 1
    		},
    		{
    			"_id" : 4,
    			"name" : "mongo-wk4:27017",
    			"health" : 1,
    			"state" : 2,
    			"stateStr" : "SECONDARY",
    			"uptime" : 24,
    			"optime" : {
    				"ts" : Timestamp(1603852985, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDurable" : {
    				"ts" : Timestamp(1603852985, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-28T02:43:05Z"),
    			"optimeDurableDate" : ISODate("2020-10-28T02:43:05Z"),
    			"lastHeartbeat" : ISODate("2020-10-28T02:43:06.799Z"),
    			"lastHeartbeatRecv" : ISODate("2020-10-28T02:43:06.307Z"),
    			"pingMs" : NumberLong(0),
    			"lastHeartbeatMessage" : "",
    			"syncSourceHost" : "mongo-wk3:27017",
    			"syncSourceId" : 3,
    			"infoMessage" : "",
    			"configVersion" : 2,
    			"configTerm" : 1
    		}
    	],
    	"ok" : 1,
    	"$clusterTime" : {
    		"clusterTime" : Timestamp(1603852985, 1),
    		"signature" : {
    			"hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
    			"keyId" : NumberLong(0)
    		}
    	},
    	"operationTime" : Timestamp(1603852985, 1)
    }



```bash
i=5
name="mongo-wk${i}"
echo ">>> container "$name

docker run --detach --name=$name --publish=270${i}:27017 \
--volume=$PWD/Workpath:/mnt/Workpath --network=mongo-net \
mongo mongod --replSet replSetI

echo "rs.addArb(\"mongo-wk${i}:27017\")"  | mongo --host $master
```

    >>> container mongo-wk5
    99c9f8a3932edf3330bfaaa8e8374bc76609e2dfa3a8ad2877fb5562df0c7b21
    MongoDB shell version v3.6.3
    connecting to: mongodb://localhost:2701/
    MongoDB server version: 4.4.1
    WARNING: shell and server versions do not match
    {
    	"ok" : 1,
    	"$clusterTime" : {
    		"clusterTime" : Timestamp(1603853110, 1),
    		"signature" : {
    			"hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
    			"keyId" : NumberLong(0)
    		}
    	},
    	"operationTime" : Timestamp(1603853110, 1)
    }
    bye



```bash
echo "rs.status();" | mongo --quiet --host $master
```

    {
    	"set" : "replSetI",
    	"date" : ISODate("2020-10-28T02:45:26.900Z"),
    	"myState" : 1,
    	"term" : NumberLong(1),
    	"syncSourceHost" : "",
    	"syncSourceId" : -1,
    	"heartbeatIntervalMillis" : NumberLong(2000),
    	"majorityVoteCount" : 3,
    	"writeMajorityCount" : 3,
    	"votingMembersCount" : 5,
    	"writableVotingMembersCount" : 4,
    	"optimes" : {
    		"lastCommittedOpTime" : {
    			"ts" : Timestamp(1603853125, 1),
    			"t" : NumberLong(1)
    		},
    		"lastCommittedWallTime" : ISODate("2020-10-28T02:45:25.586Z"),
    		"readConcernMajorityOpTime" : {
    			"ts" : Timestamp(1603853125, 1),
    			"t" : NumberLong(1)
    		},
    		"readConcernMajorityWallTime" : ISODate("2020-10-28T02:45:25.586Z"),
    		"appliedOpTime" : {
    			"ts" : Timestamp(1603853125, 1),
    			"t" : NumberLong(1)
    		},
    		"durableOpTime" : {
    			"ts" : Timestamp(1603853125, 1),
    			"t" : NumberLong(1)
    		},
    		"lastAppliedWallTime" : ISODate("2020-10-28T02:45:25.586Z"),
    		"lastDurableWallTime" : ISODate("2020-10-28T02:45:25.586Z")
    	},
    	"lastStableRecoveryTimestamp" : Timestamp(1603853095, 1),
    	"electionCandidateMetrics" : {
    		"lastElectionReason" : "electionTimeout",
    		"lastElectionDate" : ISODate("2020-10-28T02:37:55.505Z"),
    		"electionTerm" : NumberLong(1),
    		"lastCommittedOpTimeAtElection" : {
    			"ts" : Timestamp(0, 0),
    			"t" : NumberLong(-1)
    		},
    		"lastSeenOpTimeAtElection" : {
    			"ts" : Timestamp(1603852665, 1),
    			"t" : NumberLong(-1)
    		},
    		"numVotesNeeded" : 2,
    		"priorityAtElection" : 1,
    		"electionTimeoutMillis" : NumberLong(10000),
    		"numCatchUpOps" : NumberLong(0),
    		"newTermStartDate" : ISODate("2020-10-28T02:37:55.542Z"),
    		"wMajorityWriteAvailabilityDate" : ISODate("2020-10-28T02:37:57.013Z")
    	},
    	"members" : [
    		{
    			"_id" : 1,
    			"name" : "mongo-wk1:27017",
    			"health" : 1,
    			"state" : 1,
    			"stateStr" : "PRIMARY",
    			"uptime" : 498,
    			"optime" : {
    				"ts" : Timestamp(1603853125, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-28T02:45:25Z"),
    			"syncSourceHost" : "",
    			"syncSourceId" : -1,
    			"infoMessage" : "",
    			"electionTime" : Timestamp(1603852675, 1),
    			"electionDate" : ISODate("2020-10-28T02:37:55Z"),
    			"configVersion" : 3,
    			"configTerm" : 1,
    			"self" : true,
    			"lastHeartbeatMessage" : ""
    		},
    		{
    			"_id" : 2,
    			"name" : "mongo-wk2:27017",
    			"health" : 1,
    			"state" : 2,
    			"stateStr" : "SECONDARY",
    			"uptime" : 461,
    			"optime" : {
    				"ts" : Timestamp(1603853125, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDurable" : {
    				"ts" : Timestamp(1603853125, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-28T02:45:25Z"),
    			"optimeDurableDate" : ISODate("2020-10-28T02:45:25Z"),
    			"lastHeartbeat" : ISODate("2020-10-28T02:45:26.570Z"),
    			"lastHeartbeatRecv" : ISODate("2020-10-28T02:45:26.570Z"),
    			"pingMs" : NumberLong(0),
    			"lastHeartbeatMessage" : "",
    			"syncSourceHost" : "mongo-wk1:27017",
    			"syncSourceId" : 1,
    			"infoMessage" : "",
    			"configVersion" : 3,
    			"configTerm" : 1
    		},
    		{
    			"_id" : 3,
    			"name" : "mongo-wk3:27017",
    			"health" : 1,
    			"state" : 2,
    			"stateStr" : "SECONDARY",
    			"uptime" : 461,
    			"optime" : {
    				"ts" : Timestamp(1603853125, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDurable" : {
    				"ts" : Timestamp(1603853125, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-28T02:45:25Z"),
    			"optimeDurableDate" : ISODate("2020-10-28T02:45:25Z"),
    			"lastHeartbeat" : ISODate("2020-10-28T02:45:26.569Z"),
    			"lastHeartbeatRecv" : ISODate("2020-10-28T02:45:26.573Z"),
    			"pingMs" : NumberLong(0),
    			"lastHeartbeatMessage" : "",
    			"syncSourceHost" : "mongo-wk1:27017",
    			"syncSourceId" : 1,
    			"infoMessage" : "",
    			"configVersion" : 3,
    			"configTerm" : 1
    		},
    		{
    			"_id" : 4,
    			"name" : "mongo-wk4:27017",
    			"health" : 1,
    			"state" : 2,
    			"stateStr" : "SECONDARY",
    			"uptime" : 164,
    			"optime" : {
    				"ts" : Timestamp(1603853125, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDurable" : {
    				"ts" : Timestamp(1603853125, 1),
    				"t" : NumberLong(1)
    			},
    			"optimeDate" : ISODate("2020-10-28T02:45:25Z"),
    			"optimeDurableDate" : ISODate("2020-10-28T02:45:25Z"),
    			"lastHeartbeat" : ISODate("2020-10-28T02:45:26.570Z"),
    			"lastHeartbeatRecv" : ISODate("2020-10-28T02:45:26.570Z"),
    			"pingMs" : NumberLong(0),
    			"lastHeartbeatMessage" : "",
    			"syncSourceHost" : "mongo-wk3:27017",
    			"syncSourceId" : 3,
    			"infoMessage" : "",
    			"configVersion" : 3,
    			"configTerm" : 1
    		},
    		{
    			"_id" : 5,
    			"name" : "mongo-wk5:27017",
    			"health" : 1,
    			"state" : 7,
    			"stateStr" : "ARBITER",
    			"uptime" : 14,
    			"lastHeartbeat" : ISODate("2020-10-28T02:45:26.570Z"),
    			"lastHeartbeatRecv" : ISODate("2020-10-28T02:45:26.616Z"),
    			"pingMs" : NumberLong(0),
    			"lastHeartbeatMessage" : "",
    			"syncSourceHost" : "",
    			"syncSourceId" : -1,
    			"infoMessage" : "",
    			"configVersion" : 3,
    			"configTerm" : 1
    		}
    	],
    	"ok" : 1,
    	"$clusterTime" : {
    		"clusterTime" : Timestamp(1603853125, 1),
    		"signature" : {
    			"hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
    			"keyId" : NumberLong(0)
    		}
    	},
    	"operationTime" : Timestamp(1603853125, 1)
    }



```bash
docker rm -f mongo-wk3

echo "rs.remove('mongo-wk3:27017')" | mongo --host $master
```

    mongo-wk3
    MongoDB shell version v3.6.3
    connecting to: mongodb://localhost:2701/
    MongoDB server version: 4.4.1
    WARNING: shell and server versions do not match
    error: couldn't find mongo-wk3:27017 in [
    	{
    		"_id" : 1,
    		"host" : "mongo-wk1:27017",
    		"arbiterOnly" : false,
    		"buildIndexes" : true,
    		"hidden" : false,
    		"priority" : 1,
    		"tags" : {
    			
    		},
    		"slaveDelay" : NumberLong(0),
    		"votes" : 1
    	},
    	{
    		"_id" : 4,
    		"host" : "mongo-wk4:27017",
    		"arbiterOnly" : false,
    		"buildIndexes" : true,
    		"hidden" : false,
    		"priority" : 1,
    		"tags" : {
    			
    		},
    		"slaveDelay" : NumberLong(0),
    		"votes" : 1
    	},
    	{
    		"_id" : 5,
    		"host" : "mongo-wk5:27017",
    		"arbiterOnly" : true,
    		"buildIndexes" : true,
    		"hidden" : false,
    		"priority" : 0,
    		"tags" : {
    			
    		},
    		"slaveDelay" : NumberLong(0),
    		"votes" : 1
    	}
    ]
    bye

