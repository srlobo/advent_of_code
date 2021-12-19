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

res_dict = Counter()
for i in range(len(pol_tmpl) - 1):
    key = pol_tmpl[i] + pol_tmpl[i + 1]
    res_dict.update([key])

for step in range(40):
    local_res_dict = Counter()
    for key, value in res_dict.items():
        new_key1 = key[0] + pair_insertions[key]
        local_res_dict[new_key1] += value
        new_key2 = pair_insertions[key] + key[1]
        local_res_dict[new_key2] += value
    res_dict = local_res_dict

pprint(res_dict)
nc = Counter()
for key, value in res_dict.items():
    nc[key[0]] += value
    nc[key[1]] += value
nc.update(pol_tmpl[-1])

for key, value in nc.items():
    nc[key] = value // 2

print(nc)
_min = min(nc.values())
_max = max(nc.values())
print(f"{_max} - {_min} = {_max - _min}")
