import time, threading

def hello_job():
    for i in range(10):
        print("hello")
        time.sleep(1)

def run(args=[]):
    # do some stuff
    hello_thread = threading.Thread(target=hello_job, name="hello_job", args=args)
    hello_thread.start()

run()
time.sleep(12)
