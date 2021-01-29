#! /usr/bin/env python3

import os, argparse

import pandas as pd
from pandas_profiling import ProfileReport


parser = argparse.ArgumentParser(formatter_class=argparse.ArgumentDefaultsHelpFormatter)
parser.add_argument("-file", required=True, help="dataframe file")
parser.add_argument("-out", default="", help="output file, default is html")
parser.add_argument("-sep", default="\t", help="field seperator")
args = parser.parse_args()


file, out = args.file, args.out
if out == "":
    out = ".".join(file.split(".")[:-1]) + ".html"

os.makedirs(os.path.dirname(os.path.abspath(file)), exist_ok=True)


df = pd.read_csv(file, sep=args.sep)

prof = ProfileReport(df)
prof.to_file(output_file=out)
