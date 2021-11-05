import sys, json

import pandas as pd

jf = sys.argv[1]

def removesuffix(s):
    return s[:s.rindex(".")] if s.count(".") > 0 else s


results = []
ok = True

with open(jf, "r") as f:
    for line in f.readlines():
        try:
            d = json.loads(line)
            results.append(d)
        except Exception as e:
            print("{}".format(e), file=sys.stderr)
            ok = False
            break


dt = pd.DataFrame.from_records(results)
output = removesuffix(jf) + ".tsv"
dt.to_csv(output, sep="\t", index=False)

print("save to", output)
if not ok: os.exit(1)
