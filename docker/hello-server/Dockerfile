FROM centos:8

# RUN yum update -y

RUN mkdir -p /opt/bin
COPY hello.x /opt/bin

EXPOSE 8080

CMD /opt/bin/hello.x
