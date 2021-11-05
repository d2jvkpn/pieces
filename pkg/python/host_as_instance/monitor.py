import os, sys, time, json
from datetime import datetime

import psutil, GPUtil, argparse


def gpu_info() -> dict:
    gpus = GPUtil.getGPUs()
    if len(gpus) == 0: return {}
    gpu = gpus[0]

    # gpu.id, gpu.name, gpu.uuid
    return {
      "load_percent": round(gpu.load*100, 3),
      "memory_percent": round(gpu.memoryUsed/gpu.memoryTotal*100, 3),
      "temperature": gpu.temperature,
    }


def load_info(dur=1) -> dict():
    ns0, nr0 = psutil.net_io_counters().bytes_sent, psutil.net_io_counters().bytes_recv
    dr0, dw0 = psutil.disk_io_counters().read_bytes, psutil.disk_io_counters().write_bytes

    time.sleep(dur)

    nsd = (psutil.net_io_counters().bytes_sent - ns0)/dur/1024./1024.   # MB/s
    nrd = (psutil.net_io_counters().bytes_recv - nr0)/dur/1024./1024.   # MB/s
    drd = (psutil.disk_io_counters().read_bytes - dr0)/dur/1024./1024.  # MB/s
    dwd = (psutil.disk_io_counters().write_bytes - dw0)/dur/1024./1024. # MB/s

    g = gpu_info()

    return {
      "time": datetime.now().astimezone().isoformat(),
      "cpu_percent": psutil.cpu_percent(),
      "memory_percent": psutil.virtual_memory().percent,
      "swap_memory_percent": psutil.swap_memory().percent,
      "gpu_load_percent": g.get("load_percent", 0.0),
      "gpu_memory_percent": g.get("memory_percent", 0.0),
      "gpu_temperature": g.get("temperature", 0.0),
      # "network_send_MBps": round(nsd, 3),
      # "network_recv_MBps": round(nrd, 3),
      "network_send_Mbps": round(nsd*8, 3),
      "network_recv_Mbps": round(nrd*8, 3),
      "disk_read_MBps": round(drd, 3),
      "disk_write_MBps": round(dwd, 3),
    }


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Process some integers.')

    parser.add_argument('--minutes', type=float, default=1.0, help='sample minutes')
    parser.add_argument('--duration', type=float, default=1.0, help='sample duration(s)')

    parser.add_argument('--prefix', type=str, default=None,
       help='file(json, tsv) prefix to save results to',
    )

    args = parser.parse_args()

    print(">>> sample minutes: {}, duration: {}, output prefix: {}".format(
      args.minutes, args.duration, args.prefix,
    ))

    now = datetime.now()
    yes = not args.prefix is None

    if yes:
        output = "{}.{}".format(args.prefix, now.strftime("%FT%T").replace(":", "-")) + ".json"
        logF = open(output, 'w', encoding='utf8', buffering=1)
        print("save to", output)

    for _ in range(int(60*args.minutes)):
        text = json.dumps(load_info(args.duration), ensure_ascii=False)
 
        if yes:
            logF.write(text + "\n")
        else:
            print(text)

    if yes: logF.close()
    print("done")
