from multiprocessing import Pool
import subprocess

process = Pool(10)
results = []

for t in tasks:
    results.append (process.apply_async(F1, args=(a="1", b="2")))

# results = [ process.apply_async (runMT, args = (cfg, tasks, i) ) for i in objattrs]

process.close ()
process.join ()

ec = [i.get() for i in results]

####
if 1 in ec:
	print("some tasks exit with error")

out = subprocess.getoutput('whoami')

P = subprocess.Popen ("ls", 
	stdout = subprocess.PIPE, stderr=subprocess.PIPE, shell = True)

P.communicate ()
exitcode = P.returncode

# 419944026112
