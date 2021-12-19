# -*- coding: utf-8 -*-
from pprint import pprint

def calculate_fuel(x, dst):
    length = abs(dst - x)
    total = 0
    for x in range(1, length + 1):
        total += x

    return total

#with open("mini_input.txt") as f:
with open("input.txt") as f:
    pos = list(map(int, f.read().strip().split(',')))

    old_t = None
    for dst in range(min(pos), max(pos)):
        t = 0
        for x in pos:
            c = calculate_fuel(x, dst)
            t += c
        print(f"{dst} -> {t}")
        if old_t is not None and t > old_t:
            print(f"Este {old_t}")
            break
        old_t = t
