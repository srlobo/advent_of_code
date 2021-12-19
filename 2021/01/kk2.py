# -*- coding: utf-8 -*-
with open("input", "r") as f:
#with open("mini_inpu", "r") as f:
    count = 0
    previous_1 = []
    previous_2 = []
    for l in f:
        n = int(l.strip())
        if len(previous_1) < 3:
            previous_1.append(n)
            continue

        previous_1.append(n)
        print(previous_1)

        print(f"{sum(previous_1[0:3])} > {sum(previous_1[1:4])}")
        if sum(previous_1[0:3]) < sum(previous_1[1:4]):
                count += 1

        previous_1.pop(0)

print(f"Final count: {count}")
