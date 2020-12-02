FROM centos:8

ADD configs /root/configs

RUN cp /root/configs/bashrc /opt/bashrc

RUN echo "test -s /opt/bashrc && source /opt/bashrc" >> /etc/bashrc && \
    echo "test -s /opt/bashrc && source /opt/bashrc" >> /etc/profile

RUN txt="/root/configs/timezone.txt";                      \
    test -s $txt && tz=$(awk 'NR{print $1; exit}' $txt) && \
    mv /etc/localtime /etc/localtime.bk$(date +"%F") &&    \
    ln -s /usr/share/zoneinfo/$tz /etc/localtime || true


RUN yum -y update
RUN yum -y install tree pigz zip dos2unix nano vim jq wget curl git uuid
RUN yum -y install python3 python3-pip

RUN conf="/root/.config/pip/pip.conf"; mkdir -p $(dirname $conf);               \
    echo -e "[global]\nindex-url = https://pypi.douban.com/simple/" >> $conf && \
    echo -e "\n[install]\ntrusted-host = pypi.douban/simple" >> $conf 

# pip3 config set global.index-url "https://pypi.douban.com/simple/"
# pip3 config set install.trusted-host "pypi.douban/simple"
# pip3 config get install.trusted-host
# cat ~/.config/pip/pip.conf

RUN umask 022; pip3 install --no-cache --upgrade pip &&          \
    pip3 install --no-cache -r /root/configs/requirements.txt && \
    python3 -m bash_kernel.install
# jupyter kernelspec list
# jupyter kernelspec uninstall unwanted-kernel


RUN yum -y install nodejs npm &&                               \
    npm config set registry https://registry.npm.taobao.org && \
    npm install -g nodemon
# npm config list
