import os, sys, time, shlex, subprocess

import psutil, GPUtil


def gpu_info() -> dict:
    gpus = GPUtil.getGPUs()
    if len(gpus) == 0: return {}
    gpu = gpus[0]

    # gpu.id, gpu.name, gpu.uuid
    return {
      "load_percent": round(gpu.load*100, 2),
      "memory_percent": round(gpu.memoryUsed/gpu.memoryTotal*100, 2),
      "temperature": gpu.temperature,
    }


def search_process(name: str) -> [];
    procs = []
    name = name.strip()
    if name == "": return procs

    for proc in psutil.process_iter():
        # len(proc.children()) > 0
        if proc.name() == name: procs.append(proc)

   procs.sort(key=lambda p: p.create_time(), reverse=True)
   return procs


class ProcX:
    cmd, wd = [], ""
    process = None

    __cpu_interval = 1;

    def __init__(self, cmd: str):
        x = shlex.split(cmd)
        if len(x) == 0: return False
        self.cmd, self.wd = x, os.getcwd()

    def run(self) -> bool:
        if not self.process is None and self.process.is_running():
            print("!!! process is running")
            return False

        try:
            process = subprocess.Popen(self.cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
            # process.kill()
        except:
            print("!!! command start failed 1:", sys.exc_info()[0])
            return False

        print(">>> ppid:", process.pid)
        process = psutil.Process(process.pid)
        for i in range(10):
            children = process.children()
            if len(children) == 0:
                print("... wait 200 millisecond")
                time.sleep(0.2)
                continue

            self.process = children[0]
            print(">>> pid:", self.process.pid)
            break

        if self.process is None:
            print("!!! command start failed 4")
            return False

        self.n = 1
        return True

    def is_running(self) -> bool:
        if self.process is None: return False

        return self.process.is_running()

    def execute(self, action: str) -> bool:
        if self.process is None: return False

        if not self.process.is_running(): return False
        status = self.process.status()

        if status != "running" and action in ["kill", "restart", "stop"]:
            return False
        elif status != "stopped" and action == "resume":
            return False

        if action == "kill":
            self.process.kill()
        elif action == "restart":
            self.process.kill()
            self.run()
        elif action == "stop":
            self.process.suspend()
        elif action == "resume":
            self.process.resume()
        else:
            print("!!! unkonwn action:", action)
            return False

        return True

    def status(self) -> dict:
        if self.process is None: return {}

        if not self.process.is_running(): return {"status": "terminated"}

        try:
            cpu_percent = round(self.process.cpu_percent(self.__cpu_interval)/psutil.cpu_count(), 2)
        except psutil.NoSuchProcess:
            return {"status": "terminated"}
        except:
            print("Unexpected error:", sys.exc_info()[0])
            raise

        return {
            "staus": self.process.status(),
            "memory_percent": round(self.process.memory_percent(), 2),
            "cpu_percent": cpu_percent,
            "system_gpu_info": gpu_info(),
        }
