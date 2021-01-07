FROM centos:7

ADD configs /root/configs

RUN cp /root/configs/bashrc /opt/bashrc
RUN echo "test -s /opt/bashrc && source /opt/bashrc" >> /etc/bashrc && \
    echo "test -s /opt/bashrc && source /opt/bashrc" >> /etc/profile

RUN mv /etc/localtime /etc/localtime.bk && \
    ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN yum -y update && yum -y install epel-release && yum -y update

RUN yum -y install which tree pigz file zip dos2unix nano vim jq wget curl git uuid bzip2

# RUN yum -y groupinstall "Development Tools"
RUN yum -y autoremove; yum -y clean all

# htop fasd bat fd fzf ncdu httpie tmux
# taskbook goaccess

# RUN yum -y install rust cargo
# RUN cargo install exa
