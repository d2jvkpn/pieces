# jupyter-on-cloud
run jupyter server on cloud


#### 1. build docker image

```bash
docker build -t jupyter:latest .
```

#### 2. start jupyter server with container (setup by cli)

```bash
bash run.sh
```

#### 3. run jupyter server directly
```bash
jupyter notebook --generate-config -y
jupyter notebook password              # enter password

jupyter lab --ip=0.0.0.0 --no-browser --allow-root --port 9000
```

#### 4. create a self-signed certificate can be generated with openssl (*optional*)
```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
    -keyout configs/jupyter.key -out configs/jupyter.pem

jupyter lab --ip=0.0.0.0 --no-browser --allow-root --port 9000 \
    --certfile=configs/jupyter.pem --keyfile configs/jupyter.key
```
