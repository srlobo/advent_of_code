# -*- coding: utf-8 -*-
from pprint import pprint
import statistics
#with open("mini_input.txt") as f:
with open("input.txt") as f:
    pos = list(map(int, f.read().strip().split(',')))
    median = int(statistics.median(pos))
    print(f"median: {median}")
    c = 0
    for el in pos:
        c += abs(el - median)
print(c)
