#! /usr/bin/python3

import sys, json, subprocess

def kvm_live_hosts():
    result = subprocess.run(['virsh', 'list', "--name"], stdout=subprocess.PIPE)
    hosts = result.stdout.decode('utf-8').split()
    return hosts

def kvm_host_ip(name):
    result = subprocess.run(['virsh', 'domifaddr', name], stdout=subprocess.PIPE)
    resultStr = result.stdout.decode('utf-8').strip()
    ip = resultStr.split("\n")[-1].split()[-1].split("/")[0]
    # print(ip)
    return ip

hosts = [{"name": h, "ip": kvm_host_ip(h)} for h in kvm_live_hosts()]

if len(sys.argv) == 1:
    print(json.dumps(hosts))
else:
    tf = sys.argv[1]
    with open(tf, "w") as f:
        json.dump(hosts, f)
        f.write('\n')
