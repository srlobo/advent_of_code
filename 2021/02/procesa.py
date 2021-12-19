# -*- coding: utf-8 -*-
#with open("mini_input.txt") as f:
with open("input.txt") as f:
    pos = 0
    depth = 0
    for l in f:
        order, n = l.strip().split(' ')
        n = int(n)
        if order == 'forward':
            pos += n
        elif order == 'down':
            depth += n
        elif order == 'up':
            depth -= n

print(pos * depth)


