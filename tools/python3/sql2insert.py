import sys, json

import toml
import pandas as pd
import numpy as np
import pymysql


def tableHeads(cursor, table):
    cursor.execute("SHOW COLUMNS FROM `{}`;".format(table))
    dh = pd.DataFrame(cursor.fetchall())
    cols = dh.iloc[:, 0].to_list()
    return cols

def select2df(cursor, table, where = ""):
    cols = tableHeads(cursor, table)

    state = "SELECT * FROM {}".format(table)
    if where != "": state += " WHERE {}".format(where)
    print(">>> sql statement: {}".format(state))

    cursor.execute(state)
    r = cursor.fetchall()
    d = pd.DataFrame(r, columns=cols)
    return d



def r2str(record):
    def convert(e):
        # if isinstance(e, np.int64): e = int(e)
        if isinstance(e, str):
            return "'" + e + "'"
        elif e is None:
            return "NULL"
        else:
            return "{}".format(e)
   
    x = [convert(e) for e in record]
    # out = json.dumps(x, ensure_ascii=False)
    return "  (" + ", ".join(x) + ")"


# config.toml example
"""
[mysql]
host = "127.0.0.1"
port = 3306
user = "root"
password = "root"
db = "user"
charset = "utf8"
"""
(tf, table, where, prefix) = sys.argv[1:5]
# tf, table, where, prefix = "config.toml", "users", "create_timestamp >= '2021-08-31'", "users"
config = toml.load(tf)["mysql"]

charset = config.get("charset", "")
charset = "utf8" if charset == "" else charset

conn = pymysql.connect(
  host = config["host"], user = config["user"], port=config["port"],
  password = config["password"], charset = charset , db = config["db"],
)

cursor = conn.cursor()
df = select2df(cursor, table, where)
df = df.assign(**df.select_dtypes(['datetime']).astype(str).to_dict('list'))

print("    read {} records".format(df.shape[0]))

data = df.apply(r2str, axis=1).to_list()
cols = ", ".join(df.columns.to_list())

if len(data) > 0:
    stats = "INSERT INTO {}.{} ({}) VALUES\n{};\n".format(config["db"], table, cols, ",\n".join(data))

    with open(prefix + ".sql", "w", encoding="utf8") as f:
        f.write(stats)

    df.to_csv(prefix + ".tsv", sep="\t", index=False)
else:
    print("no records found")

cursor.close()
conn.commit()
conn.close()
