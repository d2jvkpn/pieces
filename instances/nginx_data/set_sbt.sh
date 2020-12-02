#! /bin/bash

# cd /opt/spark-2.4.4-bin-hadoop2.7/
# ./bin/run-example  SparkPi 100
# open in browser http://192.168.249.70:4040/jobs/

## 1. add sbt option to ~/.bashrc
echo "export SBT_OPTS=\"-Dsbt.override.build.repos=true\"" >> ~/.bashrc

## 2. set sbt repositories sopurce mirror
cp  ~/.sbt/repositories  ~/.sbt/repositories.bk
echo """[repositories]
local
huaweicloud-maven: https://repo.huaweicloud.com/repository/maven/
maven-central: https://repo1.maven.org/maven2/
sbt-plugin-repo: https://repo.scala-sbt.org/scalasbt/sbt-plugin-releases, [organization]/[module]/(scala_[scalaVersion]/)(sbt_[sbtVersion]/)[revision]/[type]s/[artifact](-[classifier]).[ext]""" > ~/.sbt/repositories


## 3. install openjdk-8-jdk spark
sudo apt install openjdk-8-jdk

## 4. create work directory
mkdir -p wk_01 && cd wk_01

## 5. 填写 build.sbt
echo """name := \"scala-spark-app\"
version := \"1.0\"
scalaVersion := \"2.11.12\"

libraryDependencies ++= Seq(
    \"org.apache.spark\" %% \"spark-core\" % \"2.4.4\"
)""" > build.sbt


## 6. sbt console
# export TERM=xterm-color
touch ScalaApp.scala
sbt console
# > import org.apache.spark.SparkContext
# > import org.apache.spark.SparkContext._
sbt run
## ignore error "ERROR ContextCleaner: Error in cleaning thread"

## 7. reference: \
#  https://www.scala-sbt.org/1.x/docs/zh-cn/Getting-Started.html
