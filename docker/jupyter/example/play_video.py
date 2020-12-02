from glob import glob

from IPython.display import Video # Audio, Image, HTML


videoList = glob("video/*.mp4")

for i in range(len(videoList)):
    print("[{}]".format(i), videoList[i])

Video(videoList[0], width=800)
