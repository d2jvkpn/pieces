FROM jupyter:latest

ADD apps /apps
RUN cat /apps/bashrc >> /opt/bashrc
RUN yum -y update

# apt -y install openjdk-11-jre openjdk-11-jre-headless
RUN yum -y install java-11-openjdk-demo.x86_64 # yum search openjdk

RUN tar -xf /apps/scala-*.tgz -C /opt &&            \
    unzip /apps/sbt-*.zip -d /opt &&                \
    mkdir -p ~/.sbt && cp /apps/repositories ~/.sbt

RUN tar -xf /apps/spark-*.tgz -C /opt

RUN pip3 install --no-cache findspark pyspark spylon-kernel
RUN python3 -m spylon_kernel install

RUN rm -r /apps
