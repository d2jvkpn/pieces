#! /usr/bin/env python3

import os, sys

import pandas as pd

# demo
# python3 df_filter.py eq_active.tsv a.tsv '(["os"] == "android") & (["app_channel"] == "huawei")'

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


tsv1, tsv2, match = sys.argv[1:4]

d1 = pd.read_csv(tsv1, sep="\t")
d2 = d1[eval(match.replace("[", "d1["))]

df2tsv(d2, tsv2, gz=(tsv2.endswith(".gz")))
