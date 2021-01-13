import os, sys, datetime

import pandas as pd
import numpy as np
import pymysql, pinyin, toml
from sqlalchemy import create_engine

tf, db, tables = sys.argv[1], sys.argv[2], sys.argv[3:]

with open(tf, "r") as f: config = toml.load(f)

print("connecting to remote database {}...".format(db))

c = config["mysql"]
conn = pymysql.connect(host = c["host"], user = c["user"], \
   password = c["password"], charset = c["charset"], db = db)

os.makedirs(db, mode=511, exist_ok=True)

cursor = conn.cursor()

if len(tables) == 0: # all tables
    cursor.execute("show tables from `{}`;".format(db))
    d = pd.DataFrame(cursor.fetchall())
    tables = d.iloc[:, 0].to_list()

createTables = "CREATE DATABASE IF NOT EXISTS {} DEFAULT CHARSET utf8;\n\n".format(db)

for table in tables:
    cursor.execute("show create table `{}`;".format(table))
    createTables += cursor.fetchall()[0][1] + ";\n\n"

    cursor.execute("select * from `{}`;".format(table))
    r = cursor.fetchall()
    df = pd.DataFrame(r)

    cursor.execute("show columns from `{}`;".format(table))
    dh = pd.DataFrame(cursor.fetchall())

    if df.shape[0] == 0:
        print("skip empty table {}".format(table))
        continue

    df.columns = dh.iloc[:, 0]
    out = "{}/{}.tsv".format(db, table)
    df.to_csv(out, sep="\t", index=False)
    print("saved {}, {}x{}".format(out, df.shape[0], df.shape[1]))


with open("{}/{}_create.sql".format(db, db), "w") as f:
     f.write(createTables)

cursor.close()
conn.close()
