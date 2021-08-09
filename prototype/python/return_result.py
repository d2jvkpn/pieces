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


r1 = div(1, 0)
print(r1.ok(), r1.data, r1.err)
print(r1)

r2 = div(1, 100)
print(r2.ok(), r2.data, r2.err)
