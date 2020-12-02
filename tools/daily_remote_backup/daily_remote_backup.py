import os, sys, time, shutil, logging, argparse
from datetime import datetime, timezone
from glob import glob

import toml, paramiko, scp, schedule


def job(c):
    client = paramiko.SSHClient()
    client.set_missing_host_key_policy(paramiko.AutoAddPolicy())

    if c.get("password", "") != "":
        client.connect(c["host"], port=c["port"], username=c["username"], password=c["password"])
    else:
        client.load_system_host_keys()
        client.connect(c["host"], port=c["port"], username=c["username"])
    # stdin, stdout, stderr = client.exec_command('ls -l')
    cp = scp.SCPClient(client.get_transport())

    p = os.path.dirname(os.path.abspath(c["copy"]["dst"]))
    os.makedirs(p, mode=511, exist_ok=True)

    logging.info(c.get("name", "connecting to server..."))
    local_time = datetime.now(timezone.utc).replace(microsecond=0).astimezone()
    now = local_time.isoformat()

    dst = c["copy"]["dst"] + "_" + now
    cp.get(c["copy"]["src"], local_path=dst, recursive=True)
    cp.close()
    client.close()

    if c.get("format", "") == "": return
    shutil.make_archive(dst, c["format"], dst)
    shutil.rmtree(dst)
    logging.info("saved {}!".format(dst+"."+c["format"]))

    n = c.get("keep", 0)
    if n <= 0: return
    fs = glob(c["copy"]["dst"] + "_*." + c["format"])
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
remote = conf["remote_backup"]
if args.once:
    retry(job, remote.get("retries", 1))(remote)
    sys.exit(1)

schedule.every().day.at(remote["clock"]).do(retry(job, remote.get("retries", 1)), remote)
# schedule.every().hour.do(job, remote)
# schedule.every(10).minutes.do(job, remote)

while True:
    schedule.run_pending()
    time.sleep(1)
