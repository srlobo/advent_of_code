# -*- coding: utf-8 -*-
with open("input", "r") as f:
    count = 0
    previous_n = None
    for l in f:
        n = int(l.strip())
        if previous_n is None:
            previous_n = n
            continue
        if n > previous_n:
            count += 1

        previous_n = n
print(count)



