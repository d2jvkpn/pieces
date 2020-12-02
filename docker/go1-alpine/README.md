# go1-alpine

#### 1 Description
Build golang docker image which has git account credential inside.


#### 2 Buid image and run container
    docker build -t go1:alpine ./

    docker run --rm -it -u=rover go1:alpine sh

#### 3 Test
Git clone a private repo which imports golang package(s) from another private repo and compile it
(go build).
