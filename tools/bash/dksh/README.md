# dksh

dksh 可以帮助基于 docker 镜像 (如 ubuntu, centos) 快速构建和运行容器, 进入 bash 交互环境, 可以方便进行数据分析、软件使用测试等。

#### Start a container workspace, usage:
```text
  dksh.sh  [OPTIONS]  <docker_image>

  Options:
    -n  <string>  set the container name, E.g. hello, if not, a random string
      (16 chars) will be assigned.

    -v  <quoted string>  mount directories to container, e.g. a:/x/a:ro, a::ro, a

    -w  <string>  set work directory replace current directory

    -u  <string>  set login user for container, default: hello
      note: without mount any directry and set environment variable by set
      empty string, namely -u 

    -U  <string>  set login exisiting user in container

    -s  <string>  set shell, default is bash

    -g  <string>  create additional groups and add user to groups

    -d <qutoted string> additional docker arguments

    -c <int> set max cpu to use

    -m <string> set max memory, followed by a suffix of m, g.
```
