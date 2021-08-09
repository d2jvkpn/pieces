```bash
docker network create redis-net
```

    08b7e83fd23cfa26a229e5a5edbba9a36e967b91f7567d0a59f02a50953ded2c



```bash
docker build -t redis:local ./
```

    Sending build context to Docker daemon  18.43kB
    Step 1/5 : FROM redis
     ---> 16ecd2772934
    Step 2/5 : ADD redis.conf /root/redis.conf
     ---> 45a8e8dc55d6
    Step 3/5 : WORKDIR /root
     ---> Running in a871c6cec0c0
    Removing intermediate container a871c6cec0c0
     ---> cef657f49ac1
    Step 4/5 : EXPOSE 6379
     ---> Running in d81f4ac2354c
    Removing intermediate container d81f4ac2354c
     ---> 9ac241b33118
    Step 5/5 : CMD ["redis-server", "/root/redis.conf"]
     ---> Running in b61a24ac09d8
    Removing intermediate container b61a24ac09d8
     ---> af61133100f3
    Successfully built af61133100f3
    Successfully tagged redis:local



```bash
wks=()
for i in $(seq 1 6); do
    # pubPort=$((7000 + $i)); --publish=$pubPort:6379
    name=redis-wk$i

	echo ">>> create container: "$name
    docker run --detach --name=$name --network=redis-net redis:local

    ip=$(docker inspect $name | jq -r '.[0].NetworkSettings.Networks."redis-net".IPAddress')
    wks+=($ip)
done

xx=$(printf ", %s" "${wks[@]}")
echo ">>> wks: ${xx:2}"
```

    >>> create container: redis-wk1
    00f4ce6759e5aa46e59d295207c73ccf259ca436e5af7f1f1a75eb806e691fe5
    >>> create container: redis-wk2
    de8e956f53175666a847bcd4d23f9624e9e1f38de1199265418fd25e1a31e2df
    >>> create container: redis-wk3
    73b67afa6f9b325b177120abea3b1a1ea7c7fd94b625ceb4c2a1d0f2e6bef461
    >>> create container: redis-wk4
    0f62967ff4f288de77ed6342739af992f2012f1d8b2795f4e3315b2f3885ce9f
    >>> create container: redis-wk5
    6c01f1f9776f69d3027c0a0784fe7390ee9d726ddb4142e94e7bac72bdcc71d3
    >>> create container: redis-wk6
    28745b90da0adad82f5c0a2b2f600537e9002b18b20860f30442cf7244d82205
    >>> wks: 172.22.0.2, 172.22.0.3, 172.22.0.4, 172.22.0.5, 172.22.0.6, 172.22.0.7



```bash
printf " %s:6379" ${wks[@]}
```

     172.22.0.2:6379 172.22.0.3:6379 172.22.0.4:6379 172.22.0.5:6379 172.22.0.6:6379 172.22.0.7:6379


```bash
echo "yes" | redis-cli --cluster create $(printf " %s:6379" ${wks[@]}) --cluster-replicas 1 --cluster-yes
# redis-cli --cluster create ${wks[@]} --cluster-replicas 1 --cluster-yes
```

    [29;1m>>> Performing hash slots allocation on 6 nodes...
    [0mMaster[0] -> Slots 0 - 5460
    Master[1] -> Slots 5461 - 10922
    Master[2] -> Slots 10923 - 16383
    Adding replica 172.22.0.6:6379 to 172.22.0.2:6379
    Adding replica 172.22.0.7:6379 to 172.22.0.3:6379
    Adding replica 172.22.0.5:6379 to 172.22.0.4:6379
    M: 8323eb5ff1f6d9f18017269f603cd5f72e92f62b 172.22.0.2:6379
       slots:[0-5460] (5461 slots) master
    M: f5b18896cda57561477216f0a1dead7fbfa5c4c9 172.22.0.3:6379
       slots:[5461-10922] (5462 slots) master
    M: 3063245b7cd4185c8d8f1ff04f804f78b880b794 172.22.0.4:6379
       slots:[10923-16383] (5461 slots) master
    S: a9e7090ba1c798d59b43f2b19c5586341e357e54 172.22.0.5:6379
       replicates 3063245b7cd4185c8d8f1ff04f804f78b880b794
    S: 827b9579b71e976ad569d59cc27f963e2a0f677f 172.22.0.6:6379
       replicates 8323eb5ff1f6d9f18017269f603cd5f72e92f62b
    S: 6d38b792ba3abc212e566e4e343dca8b78149f68 172.22.0.7:6379
       replicates f5b18896cda57561477216f0a1dead7fbfa5c4c9
    [29;1m>>> Nodes configuration updated
    [0m[29;1m>>> Assign a different config epoch to each node
    [0m[29;1m>>> Sending CLUSTER MEET messages to join the cluster
    [0mWaiting for the cluster to join
    .
    [29;1m>>> Performing Cluster Check (using node 172.22.0.2:6379)
    [0mM: 8323eb5ff1f6d9f18017269f603cd5f72e92f62b 172.22.0.2:6379
       slots:[0-5460] (5461 slots) master
       1 additional replica(s)
    S: a9e7090ba1c798d59b43f2b19c5586341e357e54 172.22.0.5:6379
       slots: (0 slots) slave
       replicates 3063245b7cd4185c8d8f1ff04f804f78b880b794
    S: 827b9579b71e976ad569d59cc27f963e2a0f677f 172.22.0.6:6379
       slots: (0 slots) slave
       replicates 8323eb5ff1f6d9f18017269f603cd5f72e92f62b
    M: 3063245b7cd4185c8d8f1ff04f804f78b880b794 172.22.0.4:6379
       slots:[10923-16383] (5461 slots) master
       1 additional replica(s)
    S: 6d38b792ba3abc212e566e4e343dca8b78149f68 172.22.0.7:6379
       slots: (0 slots) slave
       replicates f5b18896cda57561477216f0a1dead7fbfa5c4c9
    M: f5b18896cda57561477216f0a1dead7fbfa5c4c9 172.22.0.3:6379
       slots:[5461-10922] (5462 slots) master
       1 additional replica(s)
    [32;1m[OK] All nodes agree about slots configuration.
    [0m[29;1m>>> Check for open slots...
    [0m[29;1m>>> Check slots coverage...
    [0m[32;1m[OK] All 16384 slots covered.
    [0m


```bash
docker ps -a
```

    CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS               NAMES
    28745b90da0a        redis:local         "docker-entrypoint.s…"   35 seconds ago      Up 35 seconds       6379/tcp            redis-wk6
    6c01f1f9776f        redis:local         "docker-entrypoint.s…"   36 seconds ago      Up 35 seconds       6379/tcp            redis-wk5
    0f62967ff4f2        redis:local         "docker-entrypoint.s…"   36 seconds ago      Up 36 seconds       6379/tcp            redis-wk4
    73b67afa6f9b        redis:local         "docker-entrypoint.s…"   37 seconds ago      Up 36 seconds       6379/tcp            redis-wk3
    de8e956f5317        redis:local         "docker-entrypoint.s…"   37 seconds ago      Up 36 seconds       6379/tcp            redis-wk2
    00f4ce6759e5        redis:local         "docker-entrypoint.s…"   38 seconds ago      Up 37 seconds       6379/tcp            redis-wk1



```bash
echo ${wks[0]}
```

    172.22.0.2



```bash
echo ">>> run set hello world @"${wks[0]}
redis-cli -c -h ${wks[0]} set hello world
redis-cli -c -h ${wks[0]} set rover 2020

echo ">>> run get hello @"${wks[3]}
redis-cli -c -h "${wks[3]}" get hello
redis-cli -c -h "${wks[3]}" get rover
```

    >>> run set hello world @172.22.0.2
    OK
    OK
    >>> run get hello @172.22.0.5
    "world"
    "2020"



```bash
# redis-cli -c -h ${wks[0]} cluster reset
##!! redis-cli --cluster reshard ${wks[0]}:6379

redis-cli -h ${wks[0]} -p 6379 cluster nodes
```

    a9e7090ba1c798d59b43f2b19c5586341e357e54 172.22.0.5:6379@16379 slave 3063245b7cd4185c8d8f1ff04f804f78b880b794 0 1603798345659 3 connected
    827b9579b71e976ad569d59cc27f963e2a0f677f 172.22.0.6:6379@16379 slave 8323eb5ff1f6d9f18017269f603cd5f72e92f62b 0 1603798345558 1 connected
    3063245b7cd4185c8d8f1ff04f804f78b880b794 172.22.0.4:6379@16379 master - 0 1603798345056 3 connected 10923-16383
    8323eb5ff1f6d9f18017269f603cd5f72e92f62b 172.22.0.2:6379@16379 myself,master - 0 1603798345000 1 connected 0-5460
    6d38b792ba3abc212e566e4e343dca8b78149f68 172.22.0.7:6379@16379 slave f5b18896cda57561477216f0a1dead7fbfa5c4c9 0 1603798344000 2 connected
    f5b18896cda57561477216f0a1dead7fbfa5c4c9 172.22.0.3:6379@16379 master - 0 1603798345000 2 connected 5461-10922



```bash
redis-cli -c -h ${wks[0]} cluster info
```

    cluster_state:ok
    cluster_slots_assigned:16384
    cluster_slots_ok:16384
    cluster_slots_pfail:0
    cluster_slots_fail:0
    cluster_known_nodes:6
    cluster_size:3
    cluster_current_epoch:6
    cluster_my_epoch:1
    cluster_stats_messages_ping_sent:2016
    cluster_stats_messages_pong_sent:2037
    cluster_stats_messages_sent:4053
    cluster_stats_messages_ping_received:2032
    cluster_stats_messages_pong_received:2016
    cluster_stats_messages_meet_received:5
    cluster_stats_messages_received:4053



```bash
redis-cli --cluster check ${wks[0]}:6379
```

    172.22.0.2:6379 (8323eb5f...) -> 2 keys | 5461 slots | 1 slaves.
    172.22.0.4:6379 (3063245b...) -> 0 keys | 5461 slots | 1 slaves.
    172.22.0.3:6379 (f5b18896...) -> 0 keys | 5462 slots | 1 slaves.
    [32;1m[OK] 2 keys in 3 masters.
    [0m0.00 keys per slot on average.
    [29;1m>>> Performing Cluster Check (using node 172.22.0.2:6379)
    [0mM: 8323eb5ff1f6d9f18017269f603cd5f72e92f62b 172.22.0.2:6379
       slots:[0-5460] (5461 slots) master
       1 additional replica(s)
    S: a9e7090ba1c798d59b43f2b19c5586341e357e54 172.22.0.5:6379
       slots: (0 slots) slave
       replicates 3063245b7cd4185c8d8f1ff04f804f78b880b794
    S: 827b9579b71e976ad569d59cc27f963e2a0f677f 172.22.0.6:6379
       slots: (0 slots) slave
       replicates 8323eb5ff1f6d9f18017269f603cd5f72e92f62b
    M: 3063245b7cd4185c8d8f1ff04f804f78b880b794 172.22.0.4:6379
       slots:[10923-16383] (5461 slots) master
       1 additional replica(s)
    S: 6d38b792ba3abc212e566e4e343dca8b78149f68 172.22.0.7:6379
       slots: (0 slots) slave
       replicates f5b18896cda57561477216f0a1dead7fbfa5c4c9
    M: f5b18896cda57561477216f0a1dead7fbfa5c4c9 172.22.0.3:6379
       slots:[5461-10922] (5462 slots) master
       1 additional replica(s)
    [32;1m[OK] All nodes agree about slots configuration.
    [0m[29;1m>>> Check for open slots...
    [0m[29;1m>>> Check slots coverage...
    [0m[32;1m[OK] All 16384 slots covered.
    [0m


```bash
redis-cli --cluster help
```

    Cluster Manager Commands:
      create         host1:port1 ... hostN:portN
                     --cluster-replicas <arg>
      check          host:port
                     --cluster-search-multiple-owners
      info           host:port
      fix            host:port
                     --cluster-search-multiple-owners
                     --cluster-fix-with-unreachable-masters
      reshard        host:port
                     --cluster-from <arg>
                     --cluster-to <arg>
                     --cluster-slots <arg>
                     --cluster-yes
                     --cluster-timeout <arg>
                     --cluster-pipeline <arg>
                     --cluster-replace
      rebalance      host:port
                     --cluster-weight <node1=w1...nodeN=wN>
                     --cluster-use-empty-masters
                     --cluster-timeout <arg>
                     --cluster-simulate
                     --cluster-pipeline <arg>
                     --cluster-threshold <arg>
                     --cluster-replace
      add-node       new_host:new_port existing_host:existing_port
                     --cluster-slave
                     --cluster-master-id <arg>
      del-node       host:port node_id
      call           host:port command arg arg .. arg
                     --cluster-only-masters
                     --cluster-only-replicas
      set-timeout    host:port milliseconds
      import         host:port
                     --cluster-from <arg>
                     --cluster-copy
                     --cluster-replace
      backup         host:port backup_directory
      help           
    
    For check, fix, reshard, del-node, set-timeout you can specify the host and port of any working node in the cluster.
    





```bash
redis-cli --cluster reshard 172.22.0.2:6379 \
--cluster-from 8323eb5ff1f6d9f18017269f603cd5f72e92f62b \
--cluster-to 3063245b7cd4185c8d8f1ff04f804f78b880b794
```

    [29;1m>>> Performing Cluster Check (using node 172.22.0.2:6379)
    [0mM: 8323eb5ff1f6d9f18017269f603cd5f72e92f62b 172.22.0.2:6379
       slots:[0-5460] (5461 slots) master
       1 additional replica(s)
    S: a9e7090ba1c798d59b43f2b19c5586341e357e54 172.22.0.5:6379
       slots: (0 slots) slave
       replicates 3063245b7cd4185c8d8f1ff04f804f78b880b794
    S: 827b9579b71e976ad569d59cc27f963e2a0f677f 172.22.0.6:6379
       slots: (0 slots) slave
       replicates 8323eb5ff1f6d9f18017269f603cd5f72e92f62b
    M: 3063245b7cd4185c8d8f1ff04f804f78b880b794 172.22.0.4:6379
       slots:[10923-16383] (5461 slots) master
       1 additional replica(s)
    S: 6d38b792ba3abc212e566e4e343dca8b78149f68 172.22.0.7:6379
       slots: (0 slots) slave
       replicates f5b18896cda57561477216f0a1dead7fbfa5c4c9
    M: f5b18896cda57561477216f0a1dead7fbfa5c4c9 172.22.0.3:6379
       slots:[5461-10922] (5462 slots) master
       1 additional replica(s)
    [32;1m[OK] All nodes agree about slots configuration.
    [0m[29;1m>>> Check for open slots...
    [0m[29;1m>>> Check slots coverage...
    [0m[32;1m[OK] All 16384 slots covered.
    [0mHow many slots do you want to move (from 1 to 16384)? 



```bash
redis-cli --cluster check ${wks[0]}:6379
```

    172.22.0.2:6379 (8323eb5f...) -> 2 keys | 5361 slots | 1 slaves.
    172.22.0.4:6379 (3063245b...) -> 0 keys | 5561 slots | 1 slaves.
    172.22.0.3:6379 (f5b18896...) -> 0 keys | 5462 slots | 1 slaves.
    [32;1m[OK] 2 keys in 3 masters.
    [0m0.00 keys per slot on average.
    [29;1m>>> Performing Cluster Check (using node 172.22.0.2:6379)
    [0mM: 8323eb5ff1f6d9f18017269f603cd5f72e92f62b 172.22.0.2:6379
       slots:[100-5460] (5361 slots) master
       1 additional replica(s)
    S: a9e7090ba1c798d59b43f2b19c5586341e357e54 172.22.0.5:6379
       slots: (0 slots) slave
       replicates 3063245b7cd4185c8d8f1ff04f804f78b880b794
    S: 827b9579b71e976ad569d59cc27f963e2a0f677f 172.22.0.6:6379
       slots: (0 slots) slave
       replicates 8323eb5ff1f6d9f18017269f603cd5f72e92f62b
    M: 3063245b7cd4185c8d8f1ff04f804f78b880b794 172.22.0.4:6379
       slots:[0-99],[10923-16383] (5561 slots) master
       1 additional replica(s)
    S: 6d38b792ba3abc212e566e4e343dca8b78149f68 172.22.0.7:6379
       slots: (0 slots) slave
       replicates f5b18896cda57561477216f0a1dead7fbfa5c4c9
    M: f5b18896cda57561477216f0a1dead7fbfa5c4c9 172.22.0.3:6379
       slots:[5461-10922] (5462 slots) master
       1 additional replica(s)
    [32;1m[OK] All nodes agree about slots configuration.
    [0m[29;1m>>> Check for open slots...
    [0m[29;1m>>> Check slots coverage...
    [0m[32;1m[OK] All 16384 slots covered.
    [0m


```bash
redis-cli --cluster del-node ${wks[0]}:6379 6d38b792ba3abc212e566e4e343dca8b78149f68
```


```bash
redis-cli --cluster check ${wks[0]}:6379
```

    172.22.0.2:6379 (8323eb5f...) -> 2 keys | 5361 slots | 1 slaves.
    172.22.0.4:6379 (3063245b...) -> 0 keys | 5561 slots | 1 slaves.
    172.22.0.3:6379 (f5b18896...) -> 0 keys | 5462 slots | 0 slaves.
    [32;1m[OK] 2 keys in 3 masters.
    [0m0.00 keys per slot on average.
    [29;1m>>> Performing Cluster Check (using node 172.22.0.2:6379)
    [0mM: 8323eb5ff1f6d9f18017269f603cd5f72e92f62b 172.22.0.2:6379
       slots:[100-5460] (5361 slots) master
       1 additional replica(s)
    S: a9e7090ba1c798d59b43f2b19c5586341e357e54 172.22.0.5:6379
       slots: (0 slots) slave
       replicates 3063245b7cd4185c8d8f1ff04f804f78b880b794
    S: 827b9579b71e976ad569d59cc27f963e2a0f677f 172.22.0.6:6379
       slots: (0 slots) slave
       replicates 8323eb5ff1f6d9f18017269f603cd5f72e92f62b
    M: 3063245b7cd4185c8d8f1ff04f804f78b880b794 172.22.0.4:6379
       slots:[0-99],[10923-16383] (5561 slots) master
       1 additional replica(s)
    M: f5b18896cda57561477216f0a1dead7fbfa5c4c9 172.22.0.3:6379
       slots:[5461-10922] (5462 slots) master
    [32;1m[OK] All nodes agree about slots configuration.
    [0m[29;1m>>> Check for open slots...
    [0m[29;1m>>> Check slots coverage...
    [0m[32;1m[OK] All 16384 slots covered.
    [0m


```bash
name=redis-wk7

docker run --detach --name=$name --network=redis-net redis:local

ip=$(docker inspect $name | jq -r '.[0].NetworkSettings.Networks."redis-net".IPAddress')
wks+=($ip)

echo ${wks[@]}
```

    ab111d1c4c633704c6c24376765a83d5e91bd94489bbbece682aeb710a6557b3
    172.22.0.2 172.22.0.3 172.22.0.4 172.22.0.5 172.22.0.6 172.22.0.7 172.22.0.8 172.22.0.8



```bash
redis-cli --cluster add-node ${wks[-1]}:6379 ${wks[0]}:6379
```

    [29;1m>>> Adding node 172.22.0.8:6379 to cluster 172.22.0.2:6379
    [0m[29;1m>>> Performing Cluster Check (using node 172.22.0.2:6379)
    [0mM: 8323eb5ff1f6d9f18017269f603cd5f72e92f62b 172.22.0.2:6379
       slots:[100-5460] (5361 slots) master
       1 additional replica(s)
    S: a9e7090ba1c798d59b43f2b19c5586341e357e54 172.22.0.5:6379
       slots: (0 slots) slave
       replicates 3063245b7cd4185c8d8f1ff04f804f78b880b794
    S: 827b9579b71e976ad569d59cc27f963e2a0f677f 172.22.0.6:6379
       slots: (0 slots) slave
       replicates 8323eb5ff1f6d9f18017269f603cd5f72e92f62b
    M: 3063245b7cd4185c8d8f1ff04f804f78b880b794 172.22.0.4:6379
       slots:[0-99],[10923-16383] (5561 slots) master
       1 additional replica(s)
    M: f5b18896cda57561477216f0a1dead7fbfa5c4c9 172.22.0.3:6379
       slots:[5461-10922] (5462 slots) master
    [32;1m[OK] All nodes agree about slots configuration.
    [0m[29;1m>>> Check for open slots...
    [0m[29;1m>>> Check slots coverage...
    [0m[32;1m[OK] All 16384 slots covered.
    [0m[29;1m>>> Send CLUSTER MEET to node 172.22.0.8:6379 to make it join the cluster.
    [0m[32;1m[OK] New node added correctly.
    [0m


```bash
redis-cli --cluster check ${wks[0]}:6379
```

    172.22.0.2:6379 (8323eb5f...) -> 2 keys | 5361 slots | 1 slaves.
    172.22.0.4:6379 (3063245b...) -> 0 keys | 5561 slots | 1 slaves.
    172.22.0.8:6379 (3eac4a8a...) -> 0 keys | 0 slots | 0 slaves.
    172.22.0.3:6379 (f5b18896...) -> 0 keys | 5462 slots | 0 slaves.
    [32;1m[OK] 2 keys in 4 masters.
    [0m0.00 keys per slot on average.
    [29;1m>>> Performing Cluster Check (using node 172.22.0.2:6379)
    [0mM: 8323eb5ff1f6d9f18017269f603cd5f72e92f62b 172.22.0.2:6379
       slots:[100-5460] (5361 slots) master
       1 additional replica(s)
    S: a9e7090ba1c798d59b43f2b19c5586341e357e54 172.22.0.5:6379
       slots: (0 slots) slave
       replicates 3063245b7cd4185c8d8f1ff04f804f78b880b794
    S: 827b9579b71e976ad569d59cc27f963e2a0f677f 172.22.0.6:6379
       slots: (0 slots) slave
       replicates 8323eb5ff1f6d9f18017269f603cd5f72e92f62b
    M: 3063245b7cd4185c8d8f1ff04f804f78b880b794 172.22.0.4:6379
       slots:[0-99],[10923-16383] (5561 slots) master
       1 additional replica(s)
    M: 3eac4a8a954c21c24f7165d4b5f1e6637354628a 172.22.0.8:6379
       slots: (0 slots) master
    M: f5b18896cda57561477216f0a1dead7fbfa5c4c9 172.22.0.3:6379
       slots:[5461-10922] (5462 slots) master
    [32;1m[OK] All nodes agree about slots configuration.
    [0m[29;1m>>> Check for open slots...
    [0m[29;1m>>> Check slots coverage...
    [0m[32;1m[OK] All 16384 slots covered.
    [0m


```bash
redis-cli --cluster rebalance ${wks[0]}:6379
```

    [29;1m>>> Performing Cluster Check (using node 172.22.0.2:6379)
    [0m[32;1m[OK] All nodes agree about slots configuration.
    [0m[29;1m>>> Check for open slots...
    [0m[29;1m>>> Check slots coverage...
    [0m[32;1m[OK] All 16384 slots covered.
    [0m[33;1m*** No rebalancing needed! All nodes are within the 2.00% threshold.
    [0m


```bash
redis-cli --cluster check ${wks[0]}:6379
```

    172.22.0.2:6379 (8323eb5f...) -> 2 keys | 5361 slots | 1 slaves.
    172.22.0.4:6379 (3063245b...) -> 0 keys | 5561 slots | 1 slaves.
    172.22.0.8:6379 (3eac4a8a...) -> 0 keys | 0 slots | 0 slaves.
    172.22.0.3:6379 (f5b18896...) -> 0 keys | 5462 slots | 0 slaves.
    [32;1m[OK] 2 keys in 4 masters.
    [0m0.00 keys per slot on average.
    [29;1m>>> Performing Cluster Check (using node 172.22.0.2:6379)
    [0mM: 8323eb5ff1f6d9f18017269f603cd5f72e92f62b 172.22.0.2:6379
       slots:[100-5460] (5361 slots) master
       1 additional replica(s)
    S: a9e7090ba1c798d59b43f2b19c5586341e357e54 172.22.0.5:6379
       slots: (0 slots) slave
       replicates 3063245b7cd4185c8d8f1ff04f804f78b880b794
    S: 827b9579b71e976ad569d59cc27f963e2a0f677f 172.22.0.6:6379
       slots: (0 slots) slave
       replicates 8323eb5ff1f6d9f18017269f603cd5f72e92f62b
    M: 3063245b7cd4185c8d8f1ff04f804f78b880b794 172.22.0.4:6379
       slots:[0-99],[10923-16383] (5561 slots) master
       1 additional replica(s)
    M: 3eac4a8a954c21c24f7165d4b5f1e6637354628a 172.22.0.8:6379
       slots: (0 slots) master
    M: f5b18896cda57561477216f0a1dead7fbfa5c4c9 172.22.0.3:6379
       slots:[5461-10922] (5462 slots) master
    [32;1m[OK] All nodes agree about slots configuration.
    [0m[29;1m>>> Check for open slots...
    [0m[29;1m>>> Check slots coverage...
    [0m[32;1m[OK] All 16384 slots covered.
    [0m


```bash
name=redis-wk8

docker run --detach --name=$name --network=redis-net redis:local

ip=$(docker inspect $name | jq -r '.[0].NetworkSettings.Networks."redis-net".IPAddress')
wks+=($ip)

echo ${wks[@]}
```

    01060cf95285608b76d921c5fe28dc5c0c739ce457801789a26bfa759e95e413
    172.22.0.2 172.22.0.3 172.22.0.4 172.22.0.5 172.22.0.6 172.22.0.7 172.22.0.8 172.22.0.8 172.22.0.9



```bash
redis-cli --cluster add-node ${wks[-1]}:6379 ${wks[0]}:6379 --cluster-slave
# --cluster-slave --cluster-master-id <arg>  # specify the master
```

    [29;1m>>> Adding node 172.22.0.9:6379 to cluster 172.22.0.2:6379
    [0m[29;1m>>> Performing Cluster Check (using node 172.22.0.2:6379)
    [0mM: 8323eb5ff1f6d9f18017269f603cd5f72e92f62b 172.22.0.2:6379
       slots:[100-5460] (5361 slots) master
       1 additional replica(s)
    S: a9e7090ba1c798d59b43f2b19c5586341e357e54 172.22.0.5:6379
       slots: (0 slots) slave
       replicates 3063245b7cd4185c8d8f1ff04f804f78b880b794
    S: 827b9579b71e976ad569d59cc27f963e2a0f677f 172.22.0.6:6379
       slots: (0 slots) slave
       replicates 8323eb5ff1f6d9f18017269f603cd5f72e92f62b
    M: 3063245b7cd4185c8d8f1ff04f804f78b880b794 172.22.0.4:6379
       slots:[0-99],[10923-16383] (5561 slots) master
       1 additional replica(s)
    M: 3eac4a8a954c21c24f7165d4b5f1e6637354628a 172.22.0.8:6379
       slots: (0 slots) master
    M: f5b18896cda57561477216f0a1dead7fbfa5c4c9 172.22.0.3:6379
       slots:[5461-10922] (5462 slots) master
    [32;1m[OK] All nodes agree about slots configuration.
    [0m[29;1m>>> Check for open slots...
    [0m[29;1m>>> Check slots coverage...
    [0m[32;1m[OK] All 16384 slots covered.
    [0mAutomatically selected master 172.22.0.8:6379
    [29;1m>>> Send CLUSTER MEET to node 172.22.0.9:6379 to make it join the cluster.
    [0mWaiting for the cluster to join
    
    [29;1m>>> Configure node as replica of 172.22.0.8:6379.
    [0m[32;1m[OK] New node added correctly.
    [0m


```bash
redis-cli --cluster check ${wks[0]}:6379
```

    172.22.0.2:6379 (8323eb5f...) -> 2 keys | 5361 slots | 1 slaves.
    172.22.0.4:6379 (3063245b...) -> 0 keys | 5561 slots | 1 slaves.
    172.22.0.8:6379 (3eac4a8a...) -> 0 keys | 0 slots | 1 slaves.
    172.22.0.3:6379 (f5b18896...) -> 0 keys | 5462 slots | 0 slaves.
    [32;1m[OK] 2 keys in 4 masters.
    [0m0.00 keys per slot on average.
    [29;1m>>> Performing Cluster Check (using node 172.22.0.2:6379)
    [0mM: 8323eb5ff1f6d9f18017269f603cd5f72e92f62b 172.22.0.2:6379
       slots:[100-5460] (5361 slots) master
       1 additional replica(s)
    S: a9e7090ba1c798d59b43f2b19c5586341e357e54 172.22.0.5:6379
       slots: (0 slots) slave
       replicates 3063245b7cd4185c8d8f1ff04f804f78b880b794
    S: 827b9579b71e976ad569d59cc27f963e2a0f677f 172.22.0.6:6379
       slots: (0 slots) slave
       replicates 8323eb5ff1f6d9f18017269f603cd5f72e92f62b
    M: 3063245b7cd4185c8d8f1ff04f804f78b880b794 172.22.0.4:6379
       slots:[0-99],[10923-16383] (5561 slots) master
       1 additional replica(s)
    S: c8a09efebc8e9ba3ec8682c06c0437a3274303d4 172.22.0.9:6379
       slots: (0 slots) slave
       replicates 3eac4a8a954c21c24f7165d4b5f1e6637354628a
    M: 3eac4a8a954c21c24f7165d4b5f1e6637354628a 172.22.0.8:6379
       slots: (0 slots) master
       1 additional replica(s)
    M: f5b18896cda57561477216f0a1dead7fbfa5c4c9 172.22.0.3:6379
       slots:[5461-10922] (5462 slots) master
    [32;1m[OK] All nodes agree about slots configuration.
    [0m[29;1m>>> Check for open slots...
    [0m[29;1m>>> Check slots coverage...
    [0m[32;1m[OK] All 16384 slots covered.
    [0m


```bash
# redis-cli -c -h ${wks[-1]}
# >>> cluster replicate f5b18896cda57561477216f0a1dead7fbfa5c4c9

redis-cli --cluster check ${wks[0]}:6379
```

    172.22.0.2:6379 (8323eb5f...) -> 2 keys | 5361 slots | 1 slaves.
    172.22.0.4:6379 (3063245b...) -> 0 keys | 5561 slots | 1 slaves.
    172.22.0.8:6379 (3eac4a8a...) -> 0 keys | 0 slots | 0 slaves.
    172.22.0.3:6379 (f5b18896...) -> 0 keys | 5462 slots | 1 slaves.
    [32;1m[OK] 2 keys in 4 masters.
    [0m0.00 keys per slot on average.
    [29;1m>>> Performing Cluster Check (using node 172.22.0.2:6379)
    [0mM: 8323eb5ff1f6d9f18017269f603cd5f72e92f62b 172.22.0.2:6379
       slots:[100-5460] (5361 slots) master
       1 additional replica(s)
    S: a9e7090ba1c798d59b43f2b19c5586341e357e54 172.22.0.5:6379
       slots: (0 slots) slave
       replicates 3063245b7cd4185c8d8f1ff04f804f78b880b794
    S: 827b9579b71e976ad569d59cc27f963e2a0f677f 172.22.0.6:6379
       slots: (0 slots) slave
       replicates 8323eb5ff1f6d9f18017269f603cd5f72e92f62b
    M: 3063245b7cd4185c8d8f1ff04f804f78b880b794 172.22.0.4:6379
       slots:[0-99],[10923-16383] (5561 slots) master
       1 additional replica(s)
    S: c8a09efebc8e9ba3ec8682c06c0437a3274303d4 172.22.0.9:6379
       slots: (0 slots) slave
       replicates f5b18896cda57561477216f0a1dead7fbfa5c4c9
    M: 3eac4a8a954c21c24f7165d4b5f1e6637354628a 172.22.0.8:6379
       slots: (0 slots) master
    M: f5b18896cda57561477216f0a1dead7fbfa5c4c9 172.22.0.3:6379
       slots:[5461-10922] (5462 slots) master
       1 additional replica(s)
    [32;1m[OK] All nodes agree about slots configuration.
    [0m[29;1m>>> Check for open slots...
    [0m[29;1m>>> Check slots coverage...
    [0m[32;1m[OK] All 16384 slots covered.
    [0m


```bash
echo ">>> ${#wks[@]} nodes"

echo ${wks[@]}
echo ${wks[@]:3}
echo ${wks[@]:1:5}

echo "${wks[@]: -3:2}"
```

    >>> 9 nodes
    172.22.0.2 172.22.0.3 172.22.0.4 172.22.0.5 172.22.0.6 172.22.0.7 172.22.0.8 172.22.0.8 172.22.0.9
    172.22.0.5 172.22.0.6 172.22.0.7 172.22.0.8 172.22.0.8 172.22.0.9
    172.22.0.3 172.22.0.4 172.22.0.5 172.22.0.6 172.22.0.7
    172.22.0.8 172.22.0.8



```bash
redis-cli -h ${wks[1]} debug segfault
```

    Error: Server closed the connection





```bash
echo ${wks[@]}
```

    172.22.0.2 172.22.0.3 172.22.0.4 172.22.0.5 172.22.0.6 172.22.0.7 172.22.0.8 172.22.0.8 172.22.0.9



```bash
redis-cli --cluster check ${wks[2]}:6379
```

    Could not connect to Redis at 172.22.0.2:6379: No route to host
    Could not connect to Redis at 172.22.0.3:6379: No route to host
    [33;1m*** WARNING: 172.22.0.9:6379 claims to be slave of unknown node ID f5b18896cda57561477216f0a1dead7fbfa5c4c9.
    [0m[33;1m*** WARNING: 172.22.0.6:6379 claims to be slave of unknown node ID 8323eb5ff1f6d9f18017269f603cd5f72e92f62b.
    [0m172.22.0.4:6379 (3063245b...) -> 0 keys | 5561 slots | 1 slaves.
    172.22.0.8:6379 (3eac4a8a...) -> 0 keys | 0 slots | 0 slaves.
    [32;1m[OK] 0 keys in 2 masters.
    [0m0.00 keys per slot on average.
    [29;1m>>> Performing Cluster Check (using node 172.22.0.4:6379)
    [0mM: 3063245b7cd4185c8d8f1ff04f804f78b880b794 172.22.0.4:6379
       slots:[0-99],[10923-16383] (5561 slots) master
       1 additional replica(s)
    S: c8a09efebc8e9ba3ec8682c06c0437a3274303d4 172.22.0.9:6379
       slots: (0 slots) slave
       replicates f5b18896cda57561477216f0a1dead7fbfa5c4c9
    M: 3eac4a8a954c21c24f7165d4b5f1e6637354628a 172.22.0.8:6379
       slots: (0 slots) master
    S: a9e7090ba1c798d59b43f2b19c5586341e357e54 172.22.0.5:6379
       slots: (0 slots) slave
       replicates 3063245b7cd4185c8d8f1ff04f804f78b880b794
    S: 827b9579b71e976ad569d59cc27f963e2a0f677f 172.22.0.6:6379
       slots: (0 slots) slave
       replicates 8323eb5ff1f6d9f18017269f603cd5f72e92f62b
    [32;1m[OK] All nodes agree about slots configuration.
    [0m[29;1m>>> Check for open slots...
    [0m[29;1m>>> Check slots coverage...
    [0m[31;1m[ERR] Not all 16384 slots are covered by nodes.
    
    [0m




```bash
redis-cli -h ${wks[2]} -p 6379 cluster nodes
```

    8323eb5ff1f6d9f18017269f603cd5f72e92f62b 172.22.0.2:6379@16379 master,fail - 1603801324541 1603801322000 1 disconnected 100-5460
    f5b18896cda57561477216f0a1dead7fbfa5c4c9 172.22.0.3:6379@16379 master,fail - 1603801327953 1603801326949 2 connected 5461-10922
    c8a09efebc8e9ba3ec8682c06c0437a3274303d4 172.22.0.9:6379@16379 slave f5b18896cda57561477216f0a1dead7fbfa5c4c9 0 1603802002504 2 connected
    3eac4a8a954c21c24f7165d4b5f1e6637354628a 172.22.0.8:6379@16379 master - 0 1603802003608 8 connected
    3063245b7cd4185c8d8f1ff04f804f78b880b794 172.22.0.4:6379@16379 myself,master - 0 1603802002000 7 connected 0-99 10923-16383
    a9e7090ba1c798d59b43f2b19c5586341e357e54 172.22.0.5:6379@16379 slave 3063245b7cd4185c8d8f1ff04f804f78b880b794 0 1603802003507 7 connected
    827b9579b71e976ad569d59cc27f963e2a0f677f 172.22.0.6:6379@16379 slave 8323eb5ff1f6d9f18017269f603cd5f72e92f62b 0 1603802003000 1 connected



```bash
redis-cli --cluster del-node ${wks[2]}:6379 8323eb5ff1f6d9f18017269f603cd5f72e92f62b
```

    [29;1m>>> Removing node 8323eb5ff1f6d9f18017269f603cd5f72e92f62b from cluster 172.22.0.4:6379
    Could not connect to Redis at 172.22.0.2:6379: No route to host
    Could not connect to Redis at 172.22.0.3:6379: No route to host
    [0m[33;1m*** WARNING: 172.22.0.9:6379 claims to be slave of unknown node ID f5b18896cda57561477216f0a1dead7fbfa5c4c9.
    [0m[33;1m*** WARNING: 172.22.0.6:6379 claims to be slave of unknown node ID 8323eb5ff1f6d9f18017269f603cd5f72e92f62b.
    [0m[31;1m[ERR] No such node ID 8323eb5ff1f6d9f18017269f603cd5f72e92f62b
    [0m




```bash
redis-cli --cluster fix ${wks[2]}:6379
```


```bash

```

    c8a09efebc8e9ba3ec8682c06c0437a3274303d4 172.22.0.9:6379@16379 slave - 0 1603802604827 8 connected
    3eac4a8a954c21c24f7165d4b5f1e6637354628a 172.22.0.8:6379@16379 master - 0 1603802605831 8 connected
    3063245b7cd4185c8d8f1ff04f804f78b880b794 172.22.0.4:6379@16379 myself,master - 0 1603802604000 7 connected 0-99 10923-16383
    a9e7090ba1c798d59b43f2b19c5586341e357e54 172.22.0.5:6379@16379 slave 3063245b7cd4185c8d8f1ff04f804f78b880b794 0 1603802604527 7 connected
    827b9579b71e976ad569d59cc27f963e2a0f677f 172.22.0.6:6379@16379 slave - 0 1603802605000 5 connected



```bash
redis-cli -h ${wks[2]} -p 6379 cluster forget 8323eb5ff1f6d9f18017269f603cd5f72e92f62b

redis-cli -h ${wks[2]} -p 6379 cluster forget f5b18896cda57561477216f0a1dead7fbfa5c4c9
```

    (error) ERR Unknown node 8323eb5ff1f6d9f18017269f603cd5f72e92f62b
    (error) ERR Unknown node f5b18896cda57561477216f0a1dead7fbfa5c4c9



```bash
##!! remove containers

docker rm -f redis-wk{1..8}
```

    redis-wk3
    redis-wk4
    redis-wk5
    redis-wk6
    redis-wk7
    redis-wk8
    Error: No such container: redis-wk1
    Error: No such container: redis-wk2





```bash
docker ps -a
```

    CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES

