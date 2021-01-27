#! /usr/bin/env python3

import os, sys, argparse
from datetime import datetime

import pandas as pd
import pymysql, toml
from sqlalchemy import create_engine

# tf, db, tables = sys.argv[1], sys.argv[2], sys.argv[3:]
parser = argparse.ArgumentParser()
parser.add_argument("-toml", default="config.toml", help="toml config file")
parser.add_argument("-sect", default="mysql", help="mysql sect in config")
parser.add_argument("-db", required=True, help="target database")
parser.add_argument("-tables", nargs="+", required=True, help='mysql tables')
parser.add_argument("-export_df", type=bool, default =False, help="export table as data frame")
parser.add_argument("-where", default="", help="mysql select with where, work with -export_df true")
args = parser.parse_args()

tf, sect = args.toml, args.sect
db, tables = args.db, args.tables
export_df, where = args.export_df, args.where

with open(tf, "r") as f: config = toml.load(tf)

c = config[sect]
conn = pymysql.connect(host = c["host"], user = c["user"], \
   password = c["password"], charset = c["charset"], db = db)

now, joinDir = datetime.now().astimezone(), os.path.join
outdir =  "mysql_{}_{}_{}".format(db, now.strftime("%FT%H%M"), int(now.timestamp()))
os.makedirs(outdir, mode=511, exist_ok=True)


cursor = conn.cursor()

if len(tables) == 0: # all tables
    cursor.execute("show tables from `{}`;".format(db))
    d = pd.DataFrame(cursor.fetchall())
    tables = d.iloc[:, 0].to_list()

tmpl = "-- {}\n\nCREATE DATABASE IF NOT EXISTS {} DEFAULT CHARSET utf8;\n\n"
createTables = tmpl.format(now.isoformat(timespec="milliseconds"), db)

for table in tables:
    cursor.execute("show create table `{}`;".format(table))
    createTables += cursor.fetchall()[0][1] + ";\n\n"

    if not export_df: continue
    if where == "":
        state = "select * from `{}`;".format(table)
    else:
        state = "select * from `{}` where {};".format(table, where)

    print(">>> executing: {}".format(state))
    cursor.execute(state)

    r = cursor.fetchall()
    df = pd.DataFrame(r)

    cursor.execute("show columns from `{}`;".format(table))
    dh = pd.DataFrame(cursor.fetchall())

    if df.shape[0] == 0:
        print("~~~ skip empty table {}".format(table))
        continue

    df.columns = dh.iloc[:, 0]
    out = joinDir(outdir, table + ".tsv")
    df.to_csv(out, sep="\t", index=False)
    print("~~~ saved {}, {}x{}".format(out, df.shape[0], df.shape[1]))

with open(joinDir(outdir, db + "_create.sql"), "w") as f:
    f.write(createTables)

cursor.close()
conn.close()
