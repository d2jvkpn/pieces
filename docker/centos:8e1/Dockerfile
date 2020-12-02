FROM centos:8

RUN yum -y update; yum -y install epel-release; yum -y update
ADD bashrc /opt/bashrc

# ed
RUN yum -y install which nano tree dos2unix file wget git zip jq vim bzip2 pigz

RUN cp /root/.bashrc /root/.bashrc.bk$(date +'%s') && \
    echo "test -s /opt/bashrc && source /opt/bashrc" >> /root/.bashrc

RUN for i in /etc/{profile,bashrc}; do \
    cp $i $i.bk; \
    echo "test -s /opt/bashrc && source /opt/bashrc" >> $i; \
  done

# RUN yum -y groupinstall "Development Tools"
RUN yum -y autoremove; yum -y clean all
