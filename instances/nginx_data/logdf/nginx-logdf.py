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

def stripQuote(s): return s.strip("\"\'")

def toTSV(df, name, gzip=False, index=False):
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
    return  pd.read_csv(
        filename,
        encoding='gbk',
        sep='\s(?=(?:[^"]*"[^"]*")*[^"]*$)(?![^\[]*\])',
        engine='python',
        # na_values='-',
        header=None,
        usecols=[0, 2, 3, 4, 5, 6, 7, 8],
        names=['remote_addr', "remote_user",'time_local',
            'request', 'status', 'bytes', 'referer',
            'user_agent'],
        iterator=True,
        chunksize=1e5,
        converters={
            'time_local': parseDatetime,
            'request': stripQuote,
            'status': int,
            'bytes': int,
            'referer': stripQuote,
            'user_agent': stripQuote
        }
    )

if __name__ == '__main__':
    log, out = os.sys.argv[1:3]
    chunks = loadLogChunks(log)
    d1 = pd.concat(chunks)
    h1s = list(d1.columns)
    # ['remote_addr', 'remote_user', 'time_local', 'request',
    #     'status', 'bytes', 'referer', 'user_agent']

    # split cloumn request
    d2 = pd.DataFrame.from_records(d1["request"].apply(splitRequest).to_list(),
        columns=["method", "path", "query", "protocol"])

    h2s = list(d2.columns)

    d3 = pd.concat([d1.drop(["request"], axis=1), d2], axis=1)
    d3 = d3[h1s[:3] + h2s + h1s[4:]]
    toTSV(d3, out)
