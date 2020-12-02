import org.apache.spark._
import org.apache.spark.SparkContext
import org.apache.spark.SparkConf
import org.apache.log4j.{Level, Logger}

object Hello {
	def main (args: Array[String]): Unit = {
		Logger.getLogger("org.apache.spark").setLevel(Level.WARN)
		Logger.getLogger("org.eclipse.jetty.server").setLevel(Level.OFF)

		println("Hello, scala-demo!")

		val conf = new SparkConf().setMaster("local").setAppName("myApp")
		val sc = new SparkContext(conf)
		val lines = sc.textFile("build.sbt")

		println(">>> print build.sbt")
		lines.foreach(println)

		sc.stop()
	}
}
