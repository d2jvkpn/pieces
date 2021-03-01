#! /usr/bin/env bash
set -eu -o pipefail

input=$1

sed 's/ - /\t/; s/ \[/\t/; s/\] "/\t/; s/" /\t/' $input |
awk 'BEGIN{
    FS=OFS="\t";

    m=split("Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec", d, "|");
    for(i=1; i<=m; i++){months[d[i]]=sprintf("%02d",i)};

    print "remote_addr", "remote_user", "time_local", "method", "path", "query",
          "protocol", "status", "bytes", "referer", "user_agent"
}
{
    line=$0

    sub(":", " ", $3); gsub("/", " ", $3); split($3, f3, " ");
    $3=f3[3]"-"months[f3[2]]"-"f3[1]" "f3[4]" "f3[5];

    gsub(" ", "\t", $4); split($4, f4, "\t");
    if(index(f4[2], "?") == 0) {f4[2]=f4[2]"\t"} else {sub("?", "\t", f4[2])};
    $4=f4[1]"\t"f4[2]"\t"f4[3];

    sub(" ", "\t", $5); sub(" \"", "\t", $5); sub("\" ", "\t", $5)

    split($0, x, "\t");
    if (length(x) == 11) {print $0} else {print line > "/dev/stderr"}
}'




exit 0

default format:
    '$remote_addr - $remote_user [$time_local] '
    '"$request" $status $body_bytes_sent '
    '"$http_referer" "$http_user_agent"';

replace [time_local] with time_iso8601 for better record parsing

nginx:
    $request: "method", "path", "query", "protocol"
    $body_bytes_sent: bytes
    $http_user_agent: user_agent
    $http_referer: referer
    $user_agent: user_agent +++
    $request_time: The total time spent processing a request. All time values are measured in seconds with millisecond resolution
