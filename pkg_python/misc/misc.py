import os, json, datetime

def datetime2rfc3339(dt: datetime.datetime) -> str:
    return dt.astimezone().isoformat(timespec="milliseconds")

def now() -> str:
    return datetime.datetime.now().astimezone().isoformat(timespec="milliseconds")

def ts2time(ts)-> str:
    return datetime.datetime.fromtimestamp(ts).isoformat(timespec="milliseconds")


def objMarshal(obj, pretty=True) -> str:
    if pretty:
        bts = json.dumps(obj, ensure_ascii=False, indent="  ").encode('utf8')
        return bts.decode() + "\n"
    
    bts = json.dumps(obj, ensure_ascii=False).encode('utf8')
    return bts.decode()

    
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
