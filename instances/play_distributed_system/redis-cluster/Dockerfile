FROM redis

ADD redis-cluster.conf /root/redis-cluster.conf

WORKDIR /root

EXPOSE 6379
CMD ["redis-server", "/root/redis-cluster.conf"]
