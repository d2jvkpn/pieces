#! python3
# -*- coding: utf-8 -*-
import os, sys, subprocess

from notebook.auth import passwd


port, password = sys.argv[1:3]
dirpath = os.path.dirname(os.path.abspath(sys.argv[0]))


ret = subprocess.run(["jupyter", "notebook", "--generate-config", "-y"])
if ret.returncode != 0: sys.exit(1)

appendLines = """
c.NotebookApp.open_browser = False
c.NotebookApp.allow_root = True
c.NotebookApp.allow_origin = '*'
c.NotebookApp.ip = '0.0.0.0'
c.NotebookApp.port = %s
""" % port

pemFile = os.path.join(dirpath, "jupyter.pem")
keyFile = os.path.join(dirpath, "jupyter.key")

if os.path.isfile(pemFile) and os.path.isfile(keyFile):
    appendLines += ("c.NotebookApp.certfile = u'%s'\n" % pemFile)
    appendLines += ("c.NotebookApp.keyfile = u'%s'\n" % keyFile)

if password != "":
    pw = passwd(password)
    appendLines += ("c.NotebookApp.password = u'%s'\n" % pw)

script = os.path.expanduser("~/.jupyter/jupyter_notebook_config.py")
with open(script, "a") as f: f.write(appendLines)


ret = subprocess.run(["jupyter", "lab"])
if ret.returncode != 0: sys.exit(1)
