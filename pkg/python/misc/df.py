import os, json
import pandas as pd


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


def df2json(df, name, index=False)
    with open(name, 'w', encoding='utf-8') as file: 
        df.to_json(file, index=False, orient="records", indent=2, force_ascii=False)
