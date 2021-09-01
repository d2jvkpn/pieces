import sys

import toml
import pandas as pd
from sqlalchemy import create_engine

(tf, jf, table) = sys.argv[1:4]
# tf, jf, table = "config2.toml", "asset.json", "asset"

config = toml.load(tf)["mysql"]

charset = config.get("charset", "")
charset = "utf8" if charset == "" else charset

engine = create_engine('mysql+pymysql://{}:{}@{}:{}/{}'.format(
  config["user"], config["password"], config["host"], config["port"], config["db"],
))

with open(jf, 'r') as f:
    df = pd.read_json(f, orient='records')

df.to_sql(table, con=engine, if_exists="append", index=False, chunksize=1000)

engine.dispose()
