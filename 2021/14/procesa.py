# -*- coding: utf-8 -*-
from pprint import pprint
from collections import Counter

#with open("mini_input.txt") as f:
with open("input.txt") as f:
    c = 0
    pair_insertions = {}
    for line in f:
        c += 1
        line = line.strip()
        if c == 1:
            pol_tmpl = line
            continue
        if c == 2:
            continue
        try:
            key, value = line.split(' -> ')
        except:
            print(line)
            import sys
            sys.exit(1)
        pair_insertions[key] = value

for step in range(10):
    new_pol_tmpl = [pol_tmpl[0]]
    for i in range(len(pol_tmpl) - 1):
        new_pol_tmpl.append(pair_insertions[pol_tmpl[i] + pol_tmpl[i + 1]])
        new_pol_tmpl.append(pol_tmpl[i + 1])
    pol_tmpl = ''.join(new_pol_tmpl)
    print(pol_tmpl)

freq = Counter(pol_tmpl)
_min = min(freq.values())
_max = max(freq.values())
print(f"{_max} - {_min} = {_max - _min}")
