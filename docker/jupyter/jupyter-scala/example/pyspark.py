# for jupyter
import os
import findspark
os.environ["JAVA_HOME"] = "/usr/lib/jvm/java-11-openjdk-amd64/"
os.environ["SPARK_HOME"] = "/opt/spark-3.0.1-bin-hadoop3.2/"
findspark.init()

import pyspark
from pyspark.sql import SparkSession
from pyspark import SparkContext

spark = SparkSession.builder.getOrCreate()
sc = SparkContext.getOrCreate()


psrdd = sc.textFile("Dockerfile")

psrdd.first()
