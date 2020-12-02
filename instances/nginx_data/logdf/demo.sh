#! /bin/bash
set -eu -o pipefail

# using spark
spark-shell -I nginx-logdf.scala --driver-memory 4G \
  --conf spark.driver.args="data/nginx.log.gz data/nginx_log"

# using shell and awk
pigz -dc data/nginx.log.gz | sh nginx-logdf.sh -

pigz -dc data/nginx.log.gz | head -n 100 |
python3 


exit 

set nginx format:

http {
    log_format fmt1 '$remote_addr $remote_user $time_iso8601 '
    '"$request" $status $body_bytes_sent '
    '"$http_referer" "$http_user_agent" $request_time';
}
