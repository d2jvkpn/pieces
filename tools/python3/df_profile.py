#! /usr/bin/env python3

import sys

import pandas as pd
from pandas_profiling import ProfileReport

tsv, out = sys.argv[1:3]

df = pd.read_csv(tsv, sep="\t")

prof = ProfileReport(df)
prof.to_file(output_file=out)
