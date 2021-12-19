# -*- coding: utf-8 -*-
from pprint import pprint
from collections import Counter
#with open("mini_input.txt") as f:
with open("input.txt") as f:
    i = list(map(int, f.read().strip().split(',')))
    # Tail will have an structure formed with number -> repetitions
    tail_global = {}
    for day in range(256):
        print(i)
        pprint(tail_global)
        next_i = []
        tail = {}
        for n in i:
            if n == 0:
                next_i.append(6)
                if 8 not in tail:
                    tail[8] = 0
                tail[8] += 1
            else:
                next_i.append(n - 1)
        # And now tail_global
        new_tail_global = {}
        for days, reps in tail_global.items():
            if days == 0:
                if 8 not in tail:
                    tail[8] = 0
                tail[8] += reps
                if 6 not in new_tail_global:
                    new_tail_global[6] = 0
                new_tail_global[6] += reps
            else:
                d_minus_1 = days - 1
                if d_minus_1 not in new_tail_global:
                    new_tail_global[d_minus_1] = 0
                new_tail_global[d_minus_1] += reps

        for days, reps in tail.items():
            if days not in new_tail_global:
                new_tail_global[days] = 0
            new_tail_global[days] += reps

        i = next_i
        tail_global = new_tail_global
print(len(i) + sum(v for v in tail_global.values()))
print(26984457539)
