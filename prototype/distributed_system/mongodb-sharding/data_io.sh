#! /usr/bin/env bash
set -eu -o pipefail

mongostat

mongotop 30

#### backup and restore
outdir=mongo_backup$(date +"%F")
DB="myDB"
Coll="myColl"

mkdir -p $outdir

mongodump --host=localhost:27017 --forceTableScan -d $DB -o $outdir
mongodump --host=localhost:27017 --forceTableScan -d $DB --gzip -o $outdir

mongorestore --host=localhost --port=27017 --db=$DB --gzip --dir=$outdir/$DB


#### export json
at=$(date +"%F")

mongoexport --db=$DB --collection=$Coll --pretty --out=$DB.$Coll.$at.json
mongoexport --db=$DB --collection=$Coll --pretty | pigz -c > $DB.$Coll.$at.json.gz

mongoimport --db=benshar_pet --collection=shop --file=$DB.$Coll.$at.json
pigz -dc $DB.$Coll.$at.json.gz | mongoimport --host=localhost:27017 --db=$DB --collection=$Coll

#### instance
mongodump -h 127.0.0.1:27017 -d $DB -u root -p root \
  --authenticationDatabase admin  -o mongo_${DB}_backup$(date +"%F")
