# go1-alpine


build golang server program with docker image alpine


#### 1. Enable docker build --squash to minimize image size
``` bash
cat > /etc/docker/daemon.json << 'EOF'
{
  "experimental": true
}
EOF

systemctl restart docker
```
