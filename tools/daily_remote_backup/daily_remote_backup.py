import os, sys, time, shutil, logging, argparse
from datetime import datetime, timezone
from glob import glob

import toml, paramiko, scp, schedule


def scpJob(c):
    client = paramiko.SSHClient()
    client.set_missing_host_key_policy(paramiko.AutoAddPolicy())

    if c.get("password", "") != "":
        client.connect(c["host"], port=c["port"], username=c["username"], password=c["password"])
    else:
        client.load_system_host_keys()
        client.connect(c["host"], port=c["port"], username=c["username"])
    # stdin, stdout, stderr = client.exec_command('ls -l')
    cp = scp.SCPClient(client.get_transport())

    logging.info(c.get("name", "connecting to server..."))
    local_time = datetime.now(timezone.utc).replace(microsecond=0).astimezone()
    now = local_time.isoformat()

    parent = os.path.dirname(os.path.abspath(c["dst"]))
    basename = os.path.basename(c["dst"]) + "_" + now
    dst = os.path.join(parent, basename)
    os.makedirs(dst, mode=511, exist_ok=True)

    for src in c["srcList"]:
        tmp = os.path.join(dst, os.path.basename(src))
        cp.get(src, local_path=tmp, recursive=True)

    cp.close()
    client.close()

    if c.get("format", "") == "": return
    shutil.make_archive(dst, c["format"], parent, basename)
    shutil.rmtree(dst)
    logging.info("saved {}!".format(dst+"."+c["format"]))

    n = c.get("keep", 0)
    if n <= 0: return
    fs = glob(c["dst"] + "_*." + c["format"])
    fs = sorted(fs, key=os.path.getctime, reverse=True)

    for f in fs[n:]:
        logging.warning("deleting " + f)
        os.remove(f)


def retry(fn, n):
    def do(*args):
        for i in range(n):
            try:
                fn(*args)
                return
            except Exception as e:
                logging.error(fn.__name__ + ": " + str(e))
        sys.exit(1)
    return do


logging.basicConfig(
    level = logging.INFO,          
    format = '%(asctime)s %(levelname)s %(filename)s %(funcName)s[%(lineno)d]: %(message)s',
    datefmt = '%Y-%m-%dT%H:%M:%S%z',
    # filename = logFilename, filemode = 'w',
)


parser = argparse.ArgumentParser()
parser.add_argument('-toml', required=True, help='toml file')
parser.add_argument('-once', type=bool, required=False, default=False, help='run job once and exit')
args = parser.parse_args()

conf = toml.load(args.toml)
remotes = conf["remote_backup"]

if args.once:
    for remote in remotes:
        retry(scpJob, remote.get("retries", 1))(remote)
        sys.exit(1)

for remote in remotes:
    for clock in remote["clocks"]:
        print("add corn job to schedule: {} at {}".format(remote["name"], clock))
        schedule.every().day.at(clock).do(retry(scpJob, remote.get("retries", 1)), remote)
# schedule.every().hour.do(scpJob, remote)
# schedule.every(10).minutes.do(scpJob, remote)

while True:
    schedule.run_pending()
    time.sleep(1)
