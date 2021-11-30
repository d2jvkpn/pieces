import time
from datetime import datetime

# import numpy as np
from PIL import ImageGrab
# ImageGrab is macOS and Windows only

def time_tag(t):
   # datetime.now().astimezone()
   return t.strftime("%FT%T").replace(":", "-")

time.sleep(5)

for i in range(30):
    now = datetime.now();
    # screenshot = ImageGrab.grab(bbox=(250, 161, 1141, 610))
    screenshot = ImageGrab.grab()
    # img = np.array(screenshot.getdata(), np.uint8).reshape(img.size[1], img.size[0], 3)
    screenshot.save("screenshot_{}.png".format(time_tag(now), 'PNG'))
    time.sleep(1)
