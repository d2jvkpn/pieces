import sys

import toml
import pandas as pd
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


(tf, table, where, out) = sys.argv[1:5]
# tf, table, where = "config1.toml", "asset", "create_timestamp >= '2021-08-31'"
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

with open(out, "w", encoding='utf-8') as f:
    f.write(df.to_json(orient="records", force_ascii=False) + "\n")


cursor.close()
conn.commit()
conn.close()
