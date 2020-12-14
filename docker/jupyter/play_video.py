import os, sys, re
from functools import cmp_to_key
from itertools import repeat
from glob import glob

from IPython.display import Video # Audio, Image, HTML


def cmpPath(x, y):
    def mf(string):
        m = re.search("\d+\.?\d*", string)
        if m is None: return 0.0
        return float(m.group())

    fp, fb = os.path.dirname, os.path.basename
    v1 = [mf(fp(x)), mf(fb(x))]
    v2 = [mf(fp(y)), mf(fb(y))]

    if v1[0] < v2[0]:
        return -1
    elif v1[0] > v2[0]:
        return 1
    else:
        return -1 if v1[1] < v2[1] else 1


def subVideos(path="", maxLevel=0, suffix="mp4"):
   """
   maxLevel = 0 for current directory, 1 for subdirectories
   """
   result = []
   for i in range(maxLevel+1):
       x = list(repeat("*", i))
       x.insert(0, path)
       x.append("*." + suffix)
       result = glob(os.path.join(*x))
       if len(result) > 0: break

   return result


videoList = sorted(subVideos(maxLevel=1), key=cmp_to_key(cmpPath))
if len(videoList) == 0: sys.exit("no videos found")

for i in range(len(videoList)):
    print("{:4s}".format("["+str(i)+"]"), videoList[i])

Video(videoList[0], width=800)
