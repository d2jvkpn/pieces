import os, json
from datetime import datetime
from dateutil.tz import tzlocal

def getTimeTag():
    now = datetime.now()
    return "{}_{}".format(now.strftime("%F"), int(datetime.timestamp(now)))

def datetime2rfc3339(dt: datetime) -> str:
    return dt.astimezone().isoformat(timespec="milliseconds")

def now() -> str:
    return datetime.now().astimezone().isoformat(timespec="milliseconds")

def ts2time(ts)-> str:
    return datetime.fromtimestamp(ts).isoformat(timespec="milliseconds")

def utctime2Local(x):
    return x.replace(tzinfo=tzutc()).astimezone(tzlocal())

def utctime2mysql(x):
    return mongoTime2Local(x).strftime('%Y-%m-%d %H:%M:%S')

def utctime2rfc33339(x):
    return mongoTime2Local(x).strftime('%Y-%m-%dT%H:%M:%S%z')

def marshalDatetime(t): # json.dumps(dt, default=marshalDatetime)
    if isinstance(t, datetime):
        return t.isoformat(timespec="milliseconds")

def time2filetag(at) -> str:
    return "{}_{}".format(at.strftime("%FT%H%M"), int(at.timestamp()))

def objMarshal(obj, pretty=True) -> bytes:
    if pretty:
        bts = json.dumps(obj, ensure_ascii=False, indent="  ").encode('utf8')
        return bts
    
    bts = json.dumps(obj, ensure_ascii=False).encode('utf8')
    return bts

    
def logStderr(msg, *a):
    if len(a) == 0:
        print(msg, file=os.sys.stderr)
    else:
        print(msg.format(*a), file=os.sys.stderr)

        
class Result:
    data, err = None, None

    def __init__(self, data, err=None):
        self.data, self.err = data, err

    def __str__(self):
        return "data: {}\nerr: {}".format(self.data, self.err)

    def ok(self):
        return self.err is None

    def setErr(self, err):
        self.err = err


def div(a, b) -> Result:
    r = Result(0)

    if b == 0:
        r.setErr("denominator is zero")
        return r

    r.data = a/b
    return r
