#! /usr/bin/env python3

import os, sys, argparse
from datetime import datetime
from dateutil.tz import tzlocal

import pandas as pd
import pymongo, toml

def getTimeTag():
    now = datetime.now()
    return "{}_{}".format(now.strftime("%FT%H%M"), int(datetime.timestamp(now)))


def df2tsv(df, name, gz=False, index=False):
    if name == "-":
        df.to_csv(os.sys.stdout, sep="\t", index=index)
        return

    os.makedirs(os.path.dirname(os.path.abspath(name)), exist_ok=True)
    if gz:
        name = name if name.endswith(".gz") else name + ".gz" 
        df.to_csv(name, sep="\t", compression='gzip', index=index)
    else:
        df.to_csv(name, sep="\t", index=index)

    print("saved \"{}\", {} records".format(name, df.shape[0]), file=os.sys.stderr)


# tf, db, tables = sys.argv[1], sys.argv[2], sys.argv[3:]
parser = argparse.ArgumentParser(formatter_class=argparse.ArgumentDefaultsHelpFormatter)
parser.add_argument("-toml", default="config.toml", help="toml config file")
parser.add_argument("-sect", default="mongo", help="mongo sect name in config")
parser.add_argument("-db", required=True, help="target database")
parser.add_argument("-tables", nargs="+", required=True, help='mongo tables')
parser.add_argument("-export_df", type=bool, default =False, help="export table as data frame")
parser.add_argument("-where", default="", help="mongo select with where, work with -export_df true")
args = parser.parse_args()

tf, sect = args.toml, args.sect
db, tables = args.db, args.tables
export_df, where = args.export_df, args.where

if where:
    where = eval(where) # dict
else:
    where = None


with open(tf, "r") as f: config = toml.load(tf)

c = config[sect]
if c["user"] != "":
    client = pymongo.MongoClient("mongodb://{}:{}@{}:{}".format(
      c["username"], c["password"], c["host"], c["port"],
    ), tz_aware=True)
else:
   client = pymongo.MongoClient("mongodb://{}:{}".format(
      c["host"], c["port"], c["db"],
    ), tz_aware=True)


now, joinDir = datetime.now().astimezone(), os.path.join
outdir =  "mongo_{}_{}_{}".format(db, now.strftime("%FT%H%M"), int(now.timestamp()))
os.makedirs(outdir, mode=511, exist_ok=True)

for table in tables:
    if where:
        records = list(client.get_database(db).get_collection(table).find(where))
    else:
        records = list(client.get_database(db).get_collection(table).find())

    ## mongo utc time to local timezone
    for i in range(len(records)):
        for k, v in records[i].items():
            if isinstance(v, datetime): records[i][k] = v.astimezone(tzlocal())
             

    d1 = pd.DataFrame.from_records(records)
    df2tsv(d1, joinDir(outdir, "{}.tsv".format(table)))

client.close()
