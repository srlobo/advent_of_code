# -*- coding: utf-8 -*-
from pprint import pprint
with open("mini_input.txt") as f:
#with open("input.txt") as f:
    i = list(map(int, f.read().strip().split(',')))
    for day in range(80):
        next_i = []
        tail = []
        for n in i:
            if n == 0:
                next_i.append(6)
                tail.append(8)
            else:
                next_i.append(n - 1)
        i = next_i
        i.extend(tail)
print(len(i))
