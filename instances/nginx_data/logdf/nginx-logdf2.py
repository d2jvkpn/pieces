import os, pytz
from datetime import datetime
import pandas as pd

def parseDatetime(date):
    dt = datetime.strptime(date[1:-7], '%d/%b/%Y:%H:%M:%S')
    dt_tz = int(date[-6:-3]) * 60 + int(date[-3:-1])
    return dt.replace(tzinfo=pytz.FixedOffset(dt_tz))

def splitRequest(req):
    a = req.split(" ", 2)
    q = a[1].split("?", 1)
    if len(q) == 1: q = [q[0], ""]
    return [a[0]] + q + [a[2]]

def df2tsv(df, name, gzip=False, index=False):
    if name == "-":
        df.to_csv(os.sys.stdout, sep="\t", index=index)
        return

    os.makedirs(os.path.dirname(os.path.abspath(name)), exist_ok=True)
    if gzip or name.endswith(".gz"):
        name = name if name.endswith(".gz") else name + ".gz" 
        df.to_csv(name, sep="\t", compression='gzip', index=index)
    else:
        df.to_csv(name, sep="\t", index=index)

    print("saved \"{}\", {} records".format(name, df.shape[0]), file=os.sys.stderr)

def loadLogChunks (filename):
    if filename == "-": filename = os.sys.stdin
    return pd.read_csv(filename, encoding='gbk', sep=' ', engine='python',
        header=None, iterator=True, chunksize=1e5)

def parseNginx(d):
    d.iloc[:,3] = (d.iloc[:,3] + d.iloc[:,4]).apply(parseDatetime)
    d = d.drop([1, 4], axis=1)
    d.columns = ['remote_addr', "remote_user",'time_local',
        'request', 'status', 'bytes', 'referer',
        'user_agent'] + \
        list("col{}".format(i) for i in range(d.shape[1]))[8:]

    dx = pd.DataFrame.from_records(d["request"].apply(splitRequest).to_list(),
        columns=["method", "path", "query", "protocol"])

    h1s, h2s = list(d.columns), list(dx.columns)
    d1 = pd.concat([d.drop(["request"], axis=1), dx], axis=1)
    d1 = d1[h1s[:3] + h2s + h1s[4:]]
    return d1

if __name__ == '__main__':
    log, out = os.sys.argv[1:3]
    chunks = loadLogChunks(log)
    d1 = parseNginx(pd.concat(chunks))
    df2tsv(d1, out)
