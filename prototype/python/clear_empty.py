class ClearEmpty:
    '''
    清除 dict, list 中的 empty 值, 仅处理 int, float, bool, None, 以及 list, dict
    '''
    empty = {int:0, str:'', list:[], dict:{}}
    floatMin, notFalse = 1e-5, True

    def __init__(self, update = {}):
        self.empty.update(update)

    def check(self, v) -> bool:
        t = type(v)

        if v is None:
            return False
        if t is float:
            return v >= self.floatMin
        elif t is bool:
             return v or not self.notFalse
        elif t in self.empty:
            return v != self.empty[t]
        else:
            raise TypeError('unexpected type: '+ str(t))

    def clearList(self, vl) -> list:
        result = []

        for v in vl:
             if not self.check(v):
                 print(">>> clear '{}'".format(v))
                 continue

             result.append(self.clear(v) if type(v) in [list, dict] else v)

        return [v for v in result if self.check(v)]
            
    def clearDict(self, vd) -> dict:
        result = {}

        for k, v in vd.items():
             if not self.check(v):
                 print(">>> clear '{}'".format(v))
                 continue

             result[k] = self.clear(v) if type(v) in [list, dict] else v

        return dict([(k, v) for k, v in result.items() if self.check(v)])

    def clear(self, v):
        t = type(v)

        if t is list:
            return self.clearList(v)
        elif t is dict:
            return self.clearDict(v)
        else:
            raise TypeError('unexpected type: '+ str(t))


ce = ClearEmpty()

for v in [
    [1,0, "", "Hello", None],
    {"a":"hello", "b":False, "c":0, "d":[1, True, False]},
    [1,0, "", "Hello", [False, "ok"]],
    [1,0, "", "Hello", {"x":False, "y":0.01, "z":1e-6}],
]:
    print("==>", v)
    print("   ", ce.clear(v))
