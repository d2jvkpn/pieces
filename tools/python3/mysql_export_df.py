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

now = datetime.datetime.now().astimezone().isoformat(timespec="milliseconds")
joinDir, outdir = os.path.join, "mysql_" + db
os.makedirs(outdir, mode=511, exist_ok=True)


cursor = conn.cursor()

if len(tables) == 0: # all tables
    cursor.execute("show tables from `{}`;".format(db))
    d = pd.DataFrame(cursor.fetchall())
    tables = d.iloc[:, 0].to_list()

createTables = "-- {}\n\nCREATE DATABASE IF NOT EXISTS {} DEFAULT CHARSET utf8;\n\n".format(now, db)

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
    out = joinDir(outdir, table + ".tsv")
    df.to_csv(out, sep="\t", index=False)
    print("saved {}, {}x{}".format(out, df.shape[0], df.shape[1]))

with open(joinDir(outdir, db + "_create.sql"), "w") as f:
    f.write(createTables)

cursor.close()
conn.close()
