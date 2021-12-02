
# https://stackoverflow.com/questions/49005924/python-schedule-do-tasks-in-parallel/49008052

import time, datetime
from multiprocessing import Process

import schedule


def job1(x=1):
    print("{} {}".format(datetime.datetime.now().astimezone(), "job1"))

def job2():
    print("{} {}".format(datetime.datetime.now().astimezone(), "    job2"))

####
def run_schedule():
    while True:
        schedule.run_pending()
        time.sleep(1)

def run_job1_schedule():
    job1(x=1)
    schedule.every(1).seconds.do(job1, x=1)
    run_schedule()

now = datetime.datetime.now() + datetime.timedelta(seconds=10)
at = now.strftime("%H:%M:%S")
print("~~~ job2 at: {}".format(at))

def run_job2_schedule():
    schedule.every().day.at(at).do(job2)
    run_schedule()

def run_job():
    p = Process(target=run_job1_schedule)
    c = Process(target=run_job2_schedule)

    p.start()
    c.start()
    print("~~~ sleep 20s")
    time.sleep(20)

    print("~~~ p.kill(), c.terminate()")
    p.kill()
    c.terminate()

    print("~~~ sleep 5s")
    time.sleep(5)
    p.close()
    c.close()

if __name__ == "__main__":
    run_job()
