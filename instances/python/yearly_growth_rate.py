import sys, math

v1, v2, n = sys.argv[1:4]

v1, v2, n = float(v1), float(v2), int(n)

rate = 10 ** (1/n * math.log10(v2/v1))

print("Yearly growth rate: {}.".format(round(rate, 3)))
