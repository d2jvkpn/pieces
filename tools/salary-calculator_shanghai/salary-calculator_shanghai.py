import sys

import numpy as np

base = 6520 # 2022
hf = 7

if len(sys.argv) > 1:
    try:
        base = int(sys.argv[1])
    except Exception as e:
        print(e)
        sys.exit(1)

if len(sys.argv) > 2:
    try:
        hf = int(sys.argv[2])
    except Exception as e:
        print(e)
        sys.exit(1)

perc_names = ("养老保险金", "医疗保险金", "失业保险金", "工伤保险金", "生育保险金", "基本住房公积金")
perc_p = (8.0,  2.0, 0.5, 0.0,  0.0, hf)
perc_c = (16.0, 9.5, 0.5, 0.16, 1.0, hf)

ps = base * np.sum(perc_p)/100
cs = base * np.sum(perc_c)/100

print("个人税后收入: {}, 雇佣基本成本: {}".format(round(base - ps, 2), round(base + cs, 2)))
