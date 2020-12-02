FROM jupyter:latest

ADD apps /apps
RUN cat /apps/bashrc >> /opt/bashrc

ADD example /example

RUN yum -y update

# apt -y install openjdk-11-jre openjdk-11-jre-headless
RUN yum -y install java-11-openjdk-demo.x86_64 # yum search openjdk

RUN tar -xf /apps/scala-*.tgz -C /opt &&            \
    unzip /apps/sbt-*.zip -d /opt &&                \
    mkdir -p ~/.sbt && cp /apps/repositories ~/.sbt

RUN tar -xf /apps/spark-*.tgz -C /opt

RUN cd /example/scala-demo &&                           \
    /opt/sbt/bin/sbt compile && /opt/sbt/bin/sbt run && \
    rm -rf /example

# https://almond.sh/docs/quick-start-install, curl -x socks5h://localhost:1080
# RUN curl -Lo /apps/coursier.sh https://git.io/coursier-cli && \
#    sh /apps/coursier.sh launch --fork almond -- --install

RUN chmod +x /apps/coursier.sh &&                       \
    /apps/coursier.sh launch --fork almond -- --install \

RUN rm -r /apps
