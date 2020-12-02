// $ spark-shell -I nginx-logdf.scala --driver-memory 4G \
  --conf spark.driver.args="data/nginx.log.gz data/nginx_log"

import scala.io.Source
val sqlContext= new org.apache.spark.sql.SQLContext(sc)
import sqlContext.implicits._

val args = sc.getConf.get("spark.driver.args").split("\\s+")
var input = args(0)
var output = args(1)

val logFile = sc.textFile(input)

val timeFormat = new java.text.SimpleDateFormat("[dd/MMM/yyyy:HH:mm:ss Z]") // 10/Apr/2020:14:05:21 +0800
val timeFormatter = new java.text.SimpleDateFormat("yyyy-MM-dd HH:mm:ss Z")
val delimiter = """\s(?=(?:[^"]*"[^"]*")*[^"]*$)(?![^\[]*\])""".r

case class Log(ip:String, time:String, method:String, path:String,
    parameter: String, protocol:String, status:Int, size:Int,
    referer:String, user_agent:String)

// val line = """223.104.160.127 - - [10/Apr/2020:14:05:21 +0800] "GET /app/api/v2/goods/detail?ppid=578012389078&pid=2&puid=0&videoID=-1&longitude=119.1037201401745&latitude=28.532142965875288&deviceID=4AB1317A-13D8-436F-AA03-A544960A6B36&uid=0&deviceToken=b046dfd0e1ac23adcf54a8f5051e84d04a4d465bc4387ba752e4a828e3e6714d&brand=Apple&model=iPhone%207&channelName=ios HTTP/1.1" 200 5934 "-" "Cabbagebox/3.6.0 IOS""""

// 202.108.43.181 - 830c4b05350 [07/May/2020:18:43:55 +0800] "GET /api/v1/ip?type=1 HTTP/1.1" 200 46 "https://s.cabbagebox.com/recharge/crawfish?from=weibo" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36"

def parseLogLine(line: String): Log = {
    // val r = delimiter.split(line).filter(_ != "-").map(_.replaceAll("^\"|\"$", ""))
    val r = delimiter.split(line).map(_.replaceAll("^\"|\"$", ""))
    val req = r(4).split(" ")
    val uri = req(1).split("\\?", 2)
    val paramter = if (uri.length == 1) "" else uri(1)

    return Log(r(0), timeFormatter.format(timeFormat.parse(r(3))), req(0), uri(0),
        paramter, req(2), r(5).toInt, r(6).toInt,
        r(7), r(8))
}

var failedToParse = 0
def parse(line: String): Option[Log] = {
    try {
        Some(parseLogLine(line))
    } catch {
        case e: Exception => {
            println(">>> failed to parse:\n    " + line)
            failedToParse += 1
            None
        }
    }
}

val df = logFile.map(parse).filter(r => !r.isEmpty).map(_.get).toDF()

val successedToParse = df.count()
println(s"~~~ nginx log total records: $successedToParse + $failedToParse")

df.write.option("delimiter", "\t").option("header", "true").mode("overwrite").csv(output)

System.exit(0)

/*
http {
    log_format compression '$remote_addr - $remote_user [$time_local] '
                           '"$request" $status $body_bytes_sent '
                           '"$http_referer" "$http_user_agent" "$gzip_ratio"';

    server {
        gzip on;
        access_log /spool/logs/nginx-access.log compression;
        ...
    }
}

$upstream_connect_time – The time spent on establishing a connection with an upstream server
$upstream_header_time – The time between establishing a connection and receiving the first byte of the response header from the upstream server
$upstream_response_time – The time between establishing a connection and receiving the last byte of the response body from the upstream server
$request_time – The total time spent processing a request
*/
