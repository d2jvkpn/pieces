import time
from datetime import datetime

import schedule
from pytimeparse.timeparse import timeparse

def job(): 
    print(datetime.now().astimezone())

per = timeparse("10s")

schedule.every(per).seconds.do(job)

while True:
    schedule.run_pending()
    time.sleep(1)
